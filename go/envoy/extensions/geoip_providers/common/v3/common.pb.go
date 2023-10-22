// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.4
// source: envoy/extensions/geoip_providers/common/v3/common.proto

package commonv3

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

type CommonGeoipProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for geolocation headers to add to request.
	GeoHeadersToAdd *CommonGeoipProviderConfig_GeolocationHeadersToAdd `protobuf:"bytes,1,opt,name=geo_headers_to_add,json=geoHeadersToAdd,proto3" json:"geo_headers_to_add,omitempty"`
}

func (x *CommonGeoipProviderConfig) Reset() {
	*x = CommonGeoipProviderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGeoipProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGeoipProviderConfig) ProtoMessage() {}

func (x *CommonGeoipProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGeoipProviderConfig.ProtoReflect.Descriptor instead.
func (*CommonGeoipProviderConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonGeoipProviderConfig) GetGeoHeadersToAdd() *CommonGeoipProviderConfig_GeolocationHeadersToAdd {
	if x != nil {
		return x.GeoHeadersToAdd
	}
	return nil
}

// The set of geolocation headers to add to request. If any of the configured headers is present
// in the incoming request, it will be overridden by the :ref:`Geoip filter <config_http_filters_geoip>`.
// [#next-free-field: 10]
type CommonGeoipProviderConfig_GeolocationHeadersToAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the header will be used to populate the country ISO code associated with the IP address.
	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	// If set, the header will be used to populate the city associated with the IP address.
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	// If set, the header will be used to populate the region ISO code associated with the IP address.
	// The least specific subdivision will be selected as region value.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// If set, the header will be used to populate the ASN associated with the IP address.
	Asn string `protobuf:"bytes,4,opt,name=asn,proto3" json:"asn,omitempty"`
	// If set, the IP address will be checked if it belongs to any type of anonymization network (e.g. VPN, public proxy etc)
	// and header will be populated with the check result. Header value will be set to either "true" or "false" depending on the check result.
	IsAnon string `protobuf:"bytes,5,opt,name=is_anon,json=isAnon,proto3" json:"is_anon,omitempty"`
	// If set, the IP address will be checked if it belongs to a VPN and header will be populated with the check result.
	// Header value will be set to either "true" or "false" depending on the check result.
	AnonVpn string `protobuf:"bytes,6,opt,name=anon_vpn,json=anonVpn,proto3" json:"anon_vpn,omitempty"`
	// If set, the IP address will be checked if it belongs to a hosting provider and header will be populated with the check result.
	// Header value will be set to either "true" or "false" depending on the check result.
	AnonHosting string `protobuf:"bytes,7,opt,name=anon_hosting,json=anonHosting,proto3" json:"anon_hosting,omitempty"`
	// If set, the IP address will be checked if it belongs to a TOR exit node and header will be populated with the check result.
	// Header value will be set to either "true" or "false" depending on the check result.
	AnonTor string `protobuf:"bytes,8,opt,name=anon_tor,json=anonTor,proto3" json:"anon_tor,omitempty"`
	// If set, the IP address will be checked if it belongs to a public proxy and header will be populated with the check result.
	// Header value will be set to either "true" or "false" depending on the check result.
	AnonProxy string `protobuf:"bytes,9,opt,name=anon_proxy,json=anonProxy,proto3" json:"anon_proxy,omitempty"`
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) Reset() {
	*x = CommonGeoipProviderConfig_GeolocationHeadersToAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGeoipProviderConfig_GeolocationHeadersToAdd) ProtoMessage() {}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGeoipProviderConfig_GeolocationHeadersToAdd.ProtoReflect.Descriptor instead.
func (*CommonGeoipProviderConfig_GeolocationHeadersToAdd) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetAsn() string {
	if x != nil {
		return x.Asn
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetIsAnon() string {
	if x != nil {
		return x.IsAnon
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetAnonVpn() string {
	if x != nil {
		return x.AnonVpn
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetAnonHosting() string {
	if x != nil {
		return x.AnonHosting
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetAnonTor() string {
	if x != nil {
		return x.AnonTor
	}
	return ""
}

func (x *CommonGeoipProviderConfig_GeolocationHeadersToAdd) GetAnonProxy() string {
	if x != nil {
		return x.AnonProxy
	}
	return ""
}

var File_envoy_extensions_geoip_providers_common_v3_common_proto protoreflect.FileDescriptor

var file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x6f, 0x69,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x04,
	0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x6f, 0x69, 0x70, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x94, 0x01, 0x0a, 0x12,
	0x67, 0x65, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x6f, 0x69,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x6f, 0x69,
	0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0f, 0x67, 0x65, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41,
	0x64, 0x64, 0x1a, 0xf7, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x25,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01,
	0xd0, 0x01, 0x01, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x03, 0x61,
	0x73, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0,
	0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x03, 0x61, 0x73, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08,
	0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x06, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x08, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x76, 0x70, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52,
	0x07, 0x61, 0x6e, 0x6f, 0x6e, 0x56, 0x70, 0x6e, 0x12, 0x2e, 0x0a, 0x0c, 0x61, 0x6e, 0x6f, 0x6e,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x0b, 0x61, 0x6e, 0x6f,
	0x6e, 0x48, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x6e, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72,
	0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x6e, 0x54, 0x6f, 0x72,
	0x12, 0x2a, 0x0a, 0x0a, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01,
	0x01, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x42, 0xad, 0x01, 0x0a,
	0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x65, 0x6f, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescOnce sync.Once
	file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescData = file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDesc
)

func file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescGZIP() []byte {
	file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescData)
	})
	return file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDescData
}

var file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_geoip_providers_common_v3_common_proto_goTypes = []interface{}{
	(*CommonGeoipProviderConfig)(nil),                         // 0: envoy.extensions.geoip_providers.common.v3.CommonGeoipProviderConfig
	(*CommonGeoipProviderConfig_GeolocationHeadersToAdd)(nil), // 1: envoy.extensions.geoip_providers.common.v3.CommonGeoipProviderConfig.GeolocationHeadersToAdd
}
var file_envoy_extensions_geoip_providers_common_v3_common_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.geoip_providers.common.v3.CommonGeoipProviderConfig.geo_headers_to_add:type_name -> envoy.extensions.geoip_providers.common.v3.CommonGeoipProviderConfig.GeolocationHeadersToAdd
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_geoip_providers_common_v3_common_proto_init() }
func file_envoy_extensions_geoip_providers_common_v3_common_proto_init() {
	if File_envoy_extensions_geoip_providers_common_v3_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonGeoipProviderConfig); i {
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
		file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonGeoipProviderConfig_GeolocationHeadersToAdd); i {
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
			RawDescriptor: file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_geoip_providers_common_v3_common_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_geoip_providers_common_v3_common_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_geoip_providers_common_v3_common_proto_msgTypes,
	}.Build()
	File_envoy_extensions_geoip_providers_common_v3_common_proto = out.File
	file_envoy_extensions_geoip_providers_common_v3_common_proto_rawDesc = nil
	file_envoy_extensions_geoip_providers_common_v3_common_proto_goTypes = nil
	file_envoy_extensions_geoip_providers_common_v3_common_proto_depIdxs = nil
}
