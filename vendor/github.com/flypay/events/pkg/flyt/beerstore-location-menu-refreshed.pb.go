// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/beerstore-location-menu-refreshed.proto

package flyt

import (
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

type BeerstoreLocationMenuRefreshed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant *Ident `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	// Merchant ID is the LCBO Store ID
	MerchantId string `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// Location Timestamp is the time that the most recent location has been saved.
	// If the location timestamp doesn't match what's in dynamo, this event
	// will be discarded and no work will be done.
	LocationTimestamp string `protobuf:"bytes,3,opt,name=location_timestamp,json=locationTimestamp,proto3" json:"location_timestamp,omitempty"`
}

func (x *BeerstoreLocationMenuRefreshed) Reset() {
	*x = BeerstoreLocationMenuRefreshed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_beerstore_location_menu_refreshed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeerstoreLocationMenuRefreshed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeerstoreLocationMenuRefreshed) ProtoMessage() {}

func (x *BeerstoreLocationMenuRefreshed) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_beerstore_location_menu_refreshed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeerstoreLocationMenuRefreshed.ProtoReflect.Descriptor instead.
func (*BeerstoreLocationMenuRefreshed) Descriptor() ([]byte, []int) {
	return file_flyt_beerstore_location_menu_refreshed_proto_rawDescGZIP(), []int{0}
}

func (x *BeerstoreLocationMenuRefreshed) GetRestaurant() *Ident {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *BeerstoreLocationMenuRefreshed) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *BeerstoreLocationMenuRefreshed) GetLocationTimestamp() string {
	if x != nil {
		return x.LocationTimestamp
	}
	return ""
}

var File_flyt_beerstore_location_menu_refreshed_proto protoreflect.FileDescriptor

var file_flyt_beerstore_location_menu_refreshed_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6d, 0x65, 0x6e, 0x75, 0x2d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x66, 0x6c, 0x79, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x66, 0x6c, 0x79,
	0x74, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01,
	0x0a, 0x1e, 0x42, 0x65, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64,
	0x12, 0x2b, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a, 0x25, 0x82,
	0xb5, 0x18, 0x21, 0x62, 0x65, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x65, 0x64, 0x42, 0x82, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x42, 0x23, 0x42, 0x65, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58,
	0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xca, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2,
	0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_flyt_beerstore_location_menu_refreshed_proto_rawDescOnce sync.Once
	file_flyt_beerstore_location_menu_refreshed_proto_rawDescData = file_flyt_beerstore_location_menu_refreshed_proto_rawDesc
)

func file_flyt_beerstore_location_menu_refreshed_proto_rawDescGZIP() []byte {
	file_flyt_beerstore_location_menu_refreshed_proto_rawDescOnce.Do(func() {
		file_flyt_beerstore_location_menu_refreshed_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_beerstore_location_menu_refreshed_proto_rawDescData)
	})
	return file_flyt_beerstore_location_menu_refreshed_proto_rawDescData
}

var file_flyt_beerstore_location_menu_refreshed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyt_beerstore_location_menu_refreshed_proto_goTypes = []interface{}{
	(*BeerstoreLocationMenuRefreshed)(nil), // 0: flyt.BeerstoreLocationMenuRefreshed
	(*Ident)(nil),                          // 1: flyt.Ident
}
var file_flyt_beerstore_location_menu_refreshed_proto_depIdxs = []int32{
	1, // 0: flyt.BeerstoreLocationMenuRefreshed.restaurant:type_name -> flyt.Ident
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flyt_beerstore_location_menu_refreshed_proto_init() }
func file_flyt_beerstore_location_menu_refreshed_proto_init() {
	if File_flyt_beerstore_location_menu_refreshed_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	file_flyt_ident_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_beerstore_location_menu_refreshed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeerstoreLocationMenuRefreshed); i {
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
			RawDescriptor: file_flyt_beerstore_location_menu_refreshed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_beerstore_location_menu_refreshed_proto_goTypes,
		DependencyIndexes: file_flyt_beerstore_location_menu_refreshed_proto_depIdxs,
		MessageInfos:      file_flyt_beerstore_location_menu_refreshed_proto_msgTypes,
	}.Build()
	File_flyt_beerstore_location_menu_refreshed_proto = out.File
	file_flyt_beerstore_location_menu_refreshed_proto_rawDesc = nil
	file_flyt_beerstore_location_menu_refreshed_proto_goTypes = nil
	file_flyt_beerstore_location_menu_refreshed_proto_depIdxs = nil
}