// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/notification.proto

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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastRun string `protobuf:"bytes,2,opt,name=last_run,json=lastRun,proto3" json:"last_run,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_flyt_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Notification) GetLastRun() string {
	if x != nil {
		return x.LastRun
	}
	return ""
}

var File_flyt_notification_proto protoreflect.FileDescriptor

var file_flyt_notification_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x79, 0x74, 0x1a,
	0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x3a, 0x10, 0x82, 0xb5, 0x18, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x70, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x42, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58,
	0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xca, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2,
	0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_flyt_notification_proto_rawDescOnce sync.Once
	file_flyt_notification_proto_rawDescData = file_flyt_notification_proto_rawDesc
)

func file_flyt_notification_proto_rawDescGZIP() []byte {
	file_flyt_notification_proto_rawDescOnce.Do(func() {
		file_flyt_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_notification_proto_rawDescData)
	})
	return file_flyt_notification_proto_rawDescData
}

var file_flyt_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyt_notification_proto_goTypes = []interface{}{
	(*Notification)(nil), // 0: flyt.Notification
}
var file_flyt_notification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flyt_notification_proto_init() }
func file_flyt_notification_proto_init() {
	if File_flyt_notification_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_flyt_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_notification_proto_goTypes,
		DependencyIndexes: file_flyt_notification_proto_depIdxs,
		MessageInfos:      file_flyt_notification_proto_msgTypes,
	}.Build()
	File_flyt_notification_proto = out.File
	file_flyt_notification_proto_rawDesc = nil
	file_flyt_notification_proto_goTypes = nil
	file_flyt_notification_proto_depIdxs = nil
}