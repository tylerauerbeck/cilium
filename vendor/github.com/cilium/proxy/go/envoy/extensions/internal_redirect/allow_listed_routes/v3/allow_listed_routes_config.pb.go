// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/internal_redirect/allow_listed_routes/v3/allow_listed_routes_config.proto

package allow_listed_routesv3

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

// An internal redirect predicate that accepts only explicitly allowed target routes.
// [#extension: envoy.internal_redirect_predicates.allow_listed_routes]
type AllowListedRoutesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of routes that's allowed as redirect target by this predicate,
	// identified by the route's :ref:`name <envoy_v3_api_field_config.route.v3.Route.route>`.
	// Empty route names are not allowed.
	AllowedRouteNames []string `protobuf:"bytes,1,rep,name=allowed_route_names,json=allowedRouteNames,proto3" json:"allowed_route_names,omitempty"`
}

func (x *AllowListedRoutesConfig) Reset() {
	*x = AllowListedRoutesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowListedRoutesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowListedRoutesConfig) ProtoMessage() {}

func (x *AllowListedRoutesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowListedRoutesConfig.ProtoReflect.Descriptor instead.
func (*AllowListedRoutesConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescGZIP(), []int{0}
}

func (x *AllowListedRoutesConfig) GetAllowedRouteNames() []string {
	if x != nil {
		return x.AllowedRouteNames
	}
	return nil
}

var File_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x57, 0x0a, 0x17, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x13, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0xe9, 0x01, 0x0a, 0x47, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x33, 0x42, 0x1c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x2f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescData = file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDesc
)

func file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescData)
	})
	return file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDescData
}

var file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_goTypes = []interface{}{
	(*AllowListedRoutesConfig)(nil), // 0: envoy.extensions.internal_redirect.allow_listed_routes.v3.AllowListedRoutesConfig
}
var file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_init()
}
func file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_init() {
	if File_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowListedRoutesConfig); i {
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
			RawDescriptor: file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto = out.File
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_rawDesc = nil
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_goTypes = nil
	file_envoy_extensions_internal_redirect_allow_listed_routes_v3_allow_listed_routes_config_proto_depIdxs = nil
}
