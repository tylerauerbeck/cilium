// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.4
// source: envoy/config/core/v3/resolver.proto

package corev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration of DNS resolver option flags which control the behavior of the DNS resolver.
type DnsResolverOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use TCP for all DNS queries instead of the default protocol UDP.
	UseTcpForDnsLookups bool `protobuf:"varint,1,opt,name=use_tcp_for_dns_lookups,json=useTcpForDnsLookups,proto3" json:"use_tcp_for_dns_lookups,omitempty"`
	// Do not use the default search domains; only query hostnames as-is or as aliases.
	NoDefaultSearchDomain bool `protobuf:"varint,2,opt,name=no_default_search_domain,json=noDefaultSearchDomain,proto3" json:"no_default_search_domain,omitempty"`
}

func (x *DnsResolverOptions) Reset() {
	*x = DnsResolverOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_resolver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsResolverOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsResolverOptions) ProtoMessage() {}

func (x *DnsResolverOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_resolver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsResolverOptions.ProtoReflect.Descriptor instead.
func (*DnsResolverOptions) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_resolver_proto_rawDescGZIP(), []int{0}
}

func (x *DnsResolverOptions) GetUseTcpForDnsLookups() bool {
	if x != nil {
		return x.UseTcpForDnsLookups
	}
	return false
}

func (x *DnsResolverOptions) GetNoDefaultSearchDomain() bool {
	if x != nil {
		return x.NoDefaultSearchDomain
	}
	return false
}

// DNS resolution configuration which includes the underlying dns resolver addresses and options.
type DnsResolutionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of dns resolver addresses. If specified, the DNS client library will perform resolution
	// via the underlying DNS resolvers. Otherwise, the default system resolvers
	// (e.g., /etc/resolv.conf) will be used.
	Resolvers []*Address `protobuf:"bytes,1,rep,name=resolvers,proto3" json:"resolvers,omitempty"`
	// Configuration of DNS resolver option flags which control the behavior of the DNS resolver.
	DnsResolverOptions *DnsResolverOptions `protobuf:"bytes,2,opt,name=dns_resolver_options,json=dnsResolverOptions,proto3" json:"dns_resolver_options,omitempty"`
}

func (x *DnsResolutionConfig) Reset() {
	*x = DnsResolutionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_resolver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsResolutionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsResolutionConfig) ProtoMessage() {}

func (x *DnsResolutionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_resolver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsResolutionConfig.ProtoReflect.Descriptor instead.
func (*DnsResolutionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_resolver_proto_rawDescGZIP(), []int{1}
}

func (x *DnsResolutionConfig) GetResolvers() []*Address {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

func (x *DnsResolutionConfig) GetDnsResolverOptions() *DnsResolverOptions {
	if x != nil {
		return x.DnsResolverOptions
	}
	return nil
}

var File_envoy_config_core_v3_resolver_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_resolver_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x44, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34,
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x63, 0x70, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x6e,
	0x73, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x75, 0x73, 0x65, 0x54, 0x63, 0x70, 0x46, 0x6f, 0x72, 0x44, 0x6e, 0x73, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x6e, 0x6f, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6e, 0x6f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xb8, 0x01,
	0x0a, 0x13, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x14,
	0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x81, 0x01, 0x0a, 0x22, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42,
	0x0d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f,
	0x72, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_resolver_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_resolver_proto_rawDescData = file_envoy_config_core_v3_resolver_proto_rawDesc
)

func file_envoy_config_core_v3_resolver_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_resolver_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_resolver_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_resolver_proto_rawDescData)
	})
	return file_envoy_config_core_v3_resolver_proto_rawDescData
}

var file_envoy_config_core_v3_resolver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_core_v3_resolver_proto_goTypes = []interface{}{
	(*DnsResolverOptions)(nil),  // 0: envoy.config.core.v3.DnsResolverOptions
	(*DnsResolutionConfig)(nil), // 1: envoy.config.core.v3.DnsResolutionConfig
	(*Address)(nil),             // 2: envoy.config.core.v3.Address
}
var file_envoy_config_core_v3_resolver_proto_depIdxs = []int32{
	2, // 0: envoy.config.core.v3.DnsResolutionConfig.resolvers:type_name -> envoy.config.core.v3.Address
	0, // 1: envoy.config.core.v3.DnsResolutionConfig.dns_resolver_options:type_name -> envoy.config.core.v3.DnsResolverOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_resolver_proto_init() }
func file_envoy_config_core_v3_resolver_proto_init() {
	if File_envoy_config_core_v3_resolver_proto != nil {
		return
	}
	file_envoy_config_core_v3_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_resolver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsResolverOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_core_v3_resolver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsResolutionConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_resolver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_resolver_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_resolver_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_resolver_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_resolver_proto = out.File
	file_envoy_config_core_v3_resolver_proto_rawDesc = nil
	file_envoy_config_core_v3_resolver_proto_goTypes = nil
	file_envoy_config_core_v3_resolver_proto_depIdxs = nil
}
