// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package watchers

import (
	"context"
	"net"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"

	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/ipcache"
	ipcachetypes "github.com/cilium/cilium/pkg/ipcache/types"
	"github.com/cilium/cilium/pkg/k8s/resource"
	"github.com/cilium/cilium/pkg/k8s/types"
	"github.com/cilium/cilium/pkg/kvstore"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/node"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/source"
	ciliumTypes "github.com/cilium/cilium/pkg/types"
	"github.com/cilium/cilium/pkg/u8proto"
)

// initCiliumEndpointOrSlices intializes the ciliumEndpoints or ciliumEndpointSlice
func (k *K8sWatcher) initCiliumEndpointOrSlices(ctx context.Context, asyncControllers *sync.WaitGroup) {
	// If CiliumEndpointSlice feature is enabled, Cilium-agent watches CiliumEndpointSlice
	// objects instead of CiliumEndpoints. Hence, skip watching CiliumEndpoints if CiliumEndpointSlice
	// feature is enabled.
	asyncControllers.Add(1)
	if option.Config.EnableCiliumEndpointSlice {
		go k.ciliumEndpointSliceInit(ctx, asyncControllers)
	} else {
		go k.ciliumEndpointsInit(ctx, asyncControllers)
	}
}

func (k *K8sWatcher) ciliumEndpointsInit(ctx context.Context, asyncControllers *sync.WaitGroup) {
	// CiliumEndpoint objects are used for ipcache discovery until the
	// key-value store is connected
	var once sync.Once
	apiGroup := k8sAPIGroupCiliumEndpointV2

	for {
		var synced atomic.Bool
		stop := make(chan struct{})

		k.k8sResourceSynced.BlockWaitGroupToSyncResources(
			stop,
			nil,
			func() bool { return synced.Load() },
			apiGroup,
		)
		k.k8sAPIGroups.AddAPI(apiGroup)

		// Signalize that we have put node controller in the wait group to sync resources.
		once.Do(asyncControllers.Done)

		// derive another context to signal Events() in case of kvstore connection
		eventsCtx, cancel := context.WithCancel(ctx)

		go func() {
			defer close(stop)

			events := k.resources.CiliumSlimEndpoint.Events(eventsCtx)
			cache := make(map[resource.Key]*types.CiliumEndpoint)
			for event := range events {
				var err error
				switch event.Kind {
				case resource.Sync:
					synced.Store(true)
				case resource.Upsert:
					oldObj, ok := cache[event.Key]
					if !ok || !oldObj.DeepEqual(event.Object) {
						k.endpointUpdated(oldObj, event.Object)
						cache[event.Key] = event.Object
					}
				case resource.Delete:
					k.endpointDeleted(event.Object)
					delete(cache, event.Key)
				}
				event.Done(err)
			}
		}()

		select {
		case <-kvstore.Connected():
			log.Info("Connected to key-value store, stopping CiliumEndpoint watcher")
			cancel()
			k.k8sResourceSynced.CancelWaitGroupToSyncResources(apiGroup)
			k.k8sAPIGroups.RemoveAPI(apiGroup)
			<-stop
		case <-ctx.Done():
			cancel()
			<-stop
			return
		}

		select {
		case <-ctx.Done():
			return
		case <-kvstore.Client().Disconnected():
			log.Info("Disconnected from key-value store, restarting CiliumEndpoint watcher")
		}
	}
}

func (k *K8sWatcher) endpointUpdated(oldEndpoint, endpoint *types.CiliumEndpoint) {
	var namedPortsChanged bool
	defer func() {
		if namedPortsChanged {
			k.policyManager.TriggerPolicyUpdates(true, "Named ports added or updated")
		}
	}()
	var ipsAdded []string
	if oldEndpoint != nil && oldEndpoint.Networking != nil {
		// Delete the old IP addresses from the IP cache
		defer func() {
			namedPortsChanged = k.removeEndpointIPCacheMetadata(oldEndpoint, ipsAdded)
		}()
	}

	id := identity.ReservedIdentityUnmanaged
	if endpoint.Identity != nil {
		id = identity.NumericIdentity(endpoint.Identity.ID)
	}

	if endpoint.Networking == nil || endpoint.Networking.NodeIP == "" {
		// When upgrading from an older version, the nodeIP may
		// not be available yet in the CiliumEndpoint and we
		// have to wait for it to be propagated
		return
	}

	nodeIP := net.ParseIP(endpoint.Networking.NodeIP)
	if nodeIP == nil {
		log.WithField("nodeIP", endpoint.Networking.NodeIP).Warning("Unable to parse node IP while processing CiliumEndpoint update")
		return
	}

	if option.Config.EnableHighScaleIPcache &&
		!identity.IsWellKnownIdentity(id) {
		// Well-known identities are kept in the high-scale ipcache because we
		// need to be able to connect to the DNS pods to resolve FQDN policies.
		scopedLog := log.WithFields(logrus.Fields{
			logfields.Identity: id,
		})
		scopedLog.Debug("Endpoint is not well-known; skipping ipcache upsert")
		return
	}

	ips, namedPortsChanged := k.upsertEndpointIPCacheMetadata(endpoint)
	ipsAdded = append(ipsAdded, ips...)
}

func (k *K8sWatcher) endpointDeleted(endpoint *types.CiliumEndpoint) {
	namedPortsChanged := k.removeEndpointIPCacheMetadata(endpoint, nil)
	if namedPortsChanged {
		k.policyManager.TriggerPolicyUpdates(true, "Named ports deleted")
	}
}

func (k *K8sWatcher) upsertEndpointIPCacheMetadata(endpoint *types.CiliumEndpoint) ([]string, bool) {
	id := identity.ReservedIdentityUnmanaged
	if endpoint.Identity != nil {
		id = identity.NumericIdentity(endpoint.Identity.ID)
	}

	// default to the standard key
	encryptionKey := node.GetEndpointEncryptKeyIndex()
	if endpoint.Encryption != nil {
		encryptionKey = uint8(endpoint.Encryption.Key)
	}
	var ipsAdded []string

	nodeIP := net.ParseIP(endpoint.Networking.NodeIP)

	k8sMeta := &ipcachetypes.K8sMetadata{
		Namespace:  endpoint.Namespace,
		PodName:    endpoint.Name,
		NamedPorts: make(ciliumTypes.NamedPortMap, len(endpoint.NamedPorts)),
	}
	for _, port := range endpoint.NamedPorts {
		p, err := u8proto.ParseProtocol(port.Protocol)
		if err != nil {
			continue
		}
		k8sMeta.NamedPorts[port.Name] = ciliumTypes.PortProto{
			Port:  port.Port,
			Proto: uint8(p),
		}
	}

	var namedPortsChanged bool
	for _, pair := range endpoint.Networking.Addressing {
		if pair.IPV4 != "" {
			ipsAdded = append(ipsAdded, pair.IPV4)
			portsChanged, _ := k.ipcache.Upsert(pair.IPV4, nodeIP, encryptionKey, k8sMeta,
				ipcache.Identity{ID: id, Source: source.CustomResource})
			if portsChanged {
				namedPortsChanged = true
			}
		}

		if pair.IPV6 != "" {
			ipsAdded = append(ipsAdded, pair.IPV6)
			portsChanged, _ := k.ipcache.Upsert(pair.IPV6, nodeIP, encryptionKey, k8sMeta,
				ipcache.Identity{ID: id, Source: source.CustomResource})
			if portsChanged {
				namedPortsChanged = true
			}
		}
	}

	return ipsAdded, namedPortsChanged
}

func (k *K8sWatcher) removeEndpointIPCacheMetadata(endpoint *types.CiliumEndpoint, ipsAdded []string) bool {
	var namedPortsChanged bool

	for _, oldPair := range endpoint.Networking.Addressing {
		v4Added, v6Added := false, false
		for _, ipAdded := range ipsAdded {
			if ipAdded == oldPair.IPV4 {
				v4Added = true
			}
			if ipAdded == oldPair.IPV6 {
				v6Added = true
			}
		}
		if !v4Added {
			portsChanged := k.ipcache.DeleteOnMetadataMatch(oldPair.IPV4, source.CustomResource, endpoint.Namespace, endpoint.Name)
			if portsChanged {
				namedPortsChanged = true
			}
		}
		if !v6Added {
			portsChanged := k.ipcache.DeleteOnMetadataMatch(oldPair.IPV6, source.CustomResource, endpoint.Namespace, endpoint.Name)
			if portsChanged {
				namedPortsChanged = true
			}
		}
	}

	return namedPortsChanged
}
