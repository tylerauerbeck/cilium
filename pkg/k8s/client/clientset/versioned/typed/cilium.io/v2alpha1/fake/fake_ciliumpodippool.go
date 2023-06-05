// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumPodIPPools implements CiliumPodIPPoolInterface
type FakeCiliumPodIPPools struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumpodippoolsResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumpodippools"}

var ciliumpodippoolsKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2alpha1", Kind: "CiliumPodIPPool"}

// Get takes name of the ciliumPodIPPool, and returns the corresponding ciliumPodIPPool object, and an error if there is any.
func (c *FakeCiliumPodIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumPodIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumpodippoolsResource, name), &v2alpha1.CiliumPodIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumPodIPPool), err
}

// List takes label and field selectors, and returns the list of CiliumPodIPPools that match those selectors.
func (c *FakeCiliumPodIPPools) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumPodIPPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumpodippoolsResource, ciliumpodippoolsKind, opts), &v2alpha1.CiliumPodIPPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumPodIPPoolList{ListMeta: obj.(*v2alpha1.CiliumPodIPPoolList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumPodIPPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumPodIPPools.
func (c *FakeCiliumPodIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumpodippoolsResource, opts))
}

// Create takes the representation of a ciliumPodIPPool and creates it.  Returns the server's representation of the ciliumPodIPPool, and an error, if there is any.
func (c *FakeCiliumPodIPPools) Create(ctx context.Context, ciliumPodIPPool *v2alpha1.CiliumPodIPPool, opts v1.CreateOptions) (result *v2alpha1.CiliumPodIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumpodippoolsResource, ciliumPodIPPool), &v2alpha1.CiliumPodIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumPodIPPool), err
}

// Update takes the representation of a ciliumPodIPPool and updates it. Returns the server's representation of the ciliumPodIPPool, and an error, if there is any.
func (c *FakeCiliumPodIPPools) Update(ctx context.Context, ciliumPodIPPool *v2alpha1.CiliumPodIPPool, opts v1.UpdateOptions) (result *v2alpha1.CiliumPodIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumpodippoolsResource, ciliumPodIPPool), &v2alpha1.CiliumPodIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumPodIPPool), err
}

// Delete takes name of the ciliumPodIPPool and deletes it. Returns an error if one occurs.
func (c *FakeCiliumPodIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumpodippoolsResource, name, opts), &v2alpha1.CiliumPodIPPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumPodIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumpodippoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumPodIPPoolList{})
	return err
}

// Patch applies the patch and returns the patched ciliumPodIPPool.
func (c *FakeCiliumPodIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumPodIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumpodippoolsResource, name, pt, data, subresources...), &v2alpha1.CiliumPodIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumPodIPPool), err
}
