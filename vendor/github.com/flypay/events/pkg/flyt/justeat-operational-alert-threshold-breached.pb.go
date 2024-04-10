// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/justeat-operational-alert-threshold-breached.proto

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

type JustEatOperationalAlertThresholdBreached_Priority int32

const (
	JustEatOperationalAlertThresholdBreached_P1 JustEatOperationalAlertThresholdBreached_Priority = 0
	JustEatOperationalAlertThresholdBreached_P2 JustEatOperationalAlertThresholdBreached_Priority = 1
	JustEatOperationalAlertThresholdBreached_P3 JustEatOperationalAlertThresholdBreached_Priority = 2
)

// Enum value maps for JustEatOperationalAlertThresholdBreached_Priority.
var (
	JustEatOperationalAlertThresholdBreached_Priority_name = map[int32]string{
		0: "P1",
		1: "P2",
		2: "P3",
	}
	JustEatOperationalAlertThresholdBreached_Priority_value = map[string]int32{
		"P1": 0,
		"P2": 1,
		"P3": 2,
	}
)

func (x JustEatOperationalAlertThresholdBreached_Priority) Enum() *JustEatOperationalAlertThresholdBreached_Priority {
	p := new(JustEatOperationalAlertThresholdBreached_Priority)
	*p = x
	return p
}

func (x JustEatOperationalAlertThresholdBreached_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JustEatOperationalAlertThresholdBreached_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_flyt_justeat_operational_alert_threshold_breached_proto_enumTypes[0].Descriptor()
}

func (JustEatOperationalAlertThresholdBreached_Priority) Type() protoreflect.EnumType {
	return &file_flyt_justeat_operational_alert_threshold_breached_proto_enumTypes[0]
}

func (x JustEatOperationalAlertThresholdBreached_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JustEatOperationalAlertThresholdBreached_Priority.Descriptor instead.
func (JustEatOperationalAlertThresholdBreached_Priority) EnumDescriptor() ([]byte, []int) {
	return file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescGZIP(), []int{0, 0}
}

type JustEatOperationalAlertThresholdBreached struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Territory    string                                            `protobuf:"bytes,1,opt,name=territory,proto3" json:"territory,omitempty"`
	Brand        string                                            `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	TimeInterval int32                                             `protobuf:"varint,3,opt,name=time_interval,json=timeInterval,proto3" json:"time_interval,omitempty"`
	Threshold    int32                                             `protobuf:"varint,4,opt,name=threshold,proto3" json:"threshold,omitempty"`
	NumFailed    int32                                             `protobuf:"varint,5,opt,name=num_failed,json=numFailed,proto3" json:"num_failed,omitempty"`
	Priority     JustEatOperationalAlertThresholdBreached_Priority `protobuf:"varint,6,opt,name=priority,proto3,enum=flyt.JustEatOperationalAlertThresholdBreached_Priority" json:"priority,omitempty"`
	Runbook      string                                            `protobuf:"bytes,7,opt,name=runbook,proto3" json:"runbook,omitempty"`
}

func (x *JustEatOperationalAlertThresholdBreached) Reset() {
	*x = JustEatOperationalAlertThresholdBreached{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_justeat_operational_alert_threshold_breached_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JustEatOperationalAlertThresholdBreached) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JustEatOperationalAlertThresholdBreached) ProtoMessage() {}

func (x *JustEatOperationalAlertThresholdBreached) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_justeat_operational_alert_threshold_breached_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JustEatOperationalAlertThresholdBreached.ProtoReflect.Descriptor instead.
func (*JustEatOperationalAlertThresholdBreached) Descriptor() ([]byte, []int) {
	return file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescGZIP(), []int{0}
}

func (x *JustEatOperationalAlertThresholdBreached) GetTerritory() string {
	if x != nil {
		return x.Territory
	}
	return ""
}

func (x *JustEatOperationalAlertThresholdBreached) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *JustEatOperationalAlertThresholdBreached) GetTimeInterval() int32 {
	if x != nil {
		return x.TimeInterval
	}
	return 0
}

func (x *JustEatOperationalAlertThresholdBreached) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *JustEatOperationalAlertThresholdBreached) GetNumFailed() int32 {
	if x != nil {
		return x.NumFailed
	}
	return 0
}

func (x *JustEatOperationalAlertThresholdBreached) GetPriority() JustEatOperationalAlertThresholdBreached_Priority {
	if x != nil {
		return x.Priority
	}
	return JustEatOperationalAlertThresholdBreached_P1
}

func (x *JustEatOperationalAlertThresholdBreached) GetRunbook() string {
	if x != nil {
		return x.Runbook
	}
	return ""
}

var File_flyt_justeat_operational_alert_threshold_breached_proto protoreflect.FileDescriptor

var file_flyt_justeat_operational_alert_threshold_breached_proto_rawDesc = []byte{
	0x0a, 0x37, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x2d, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2d, 0x62, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x79, 0x74, 0x1a,
	0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a, 0x28, 0x4a, 0x75, 0x73, 0x74, 0x45,
	0x61, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75,
	0x6d, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6e, 0x75, 0x6d, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x2e, 0x4a, 0x75, 0x73, 0x74, 0x45, 0x61, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x42, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x75, 0x6e, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x75, 0x6e, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x22, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x50, 0x32, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x33, 0x10, 0x02, 0x3a, 0x30, 0x82, 0xb5,
	0x18, 0x2c, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2e, 0x62, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x8c,
	0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x42, 0x2d, 0x4a, 0x75, 0x73,
	0x74, 0x65, 0x61, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0xa2,
	0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xca, 0x02, 0x04, 0x46,
	0x6c, 0x79, 0x74, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescOnce sync.Once
	file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescData = file_flyt_justeat_operational_alert_threshold_breached_proto_rawDesc
)

func file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescGZIP() []byte {
	file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescOnce.Do(func() {
		file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescData)
	})
	return file_flyt_justeat_operational_alert_threshold_breached_proto_rawDescData
}

var file_flyt_justeat_operational_alert_threshold_breached_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyt_justeat_operational_alert_threshold_breached_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyt_justeat_operational_alert_threshold_breached_proto_goTypes = []interface{}{
	(JustEatOperationalAlertThresholdBreached_Priority)(0), // 0: flyt.JustEatOperationalAlertThresholdBreached.Priority
	(*JustEatOperationalAlertThresholdBreached)(nil),       // 1: flyt.JustEatOperationalAlertThresholdBreached
}
var file_flyt_justeat_operational_alert_threshold_breached_proto_depIdxs = []int32{
	0, // 0: flyt.JustEatOperationalAlertThresholdBreached.priority:type_name -> flyt.JustEatOperationalAlertThresholdBreached.Priority
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flyt_justeat_operational_alert_threshold_breached_proto_init() }
func file_flyt_justeat_operational_alert_threshold_breached_proto_init() {
	if File_flyt_justeat_operational_alert_threshold_breached_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_justeat_operational_alert_threshold_breached_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JustEatOperationalAlertThresholdBreached); i {
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
			RawDescriptor: file_flyt_justeat_operational_alert_threshold_breached_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_justeat_operational_alert_threshold_breached_proto_goTypes,
		DependencyIndexes: file_flyt_justeat_operational_alert_threshold_breached_proto_depIdxs,
		EnumInfos:         file_flyt_justeat_operational_alert_threshold_breached_proto_enumTypes,
		MessageInfos:      file_flyt_justeat_operational_alert_threshold_breached_proto_msgTypes,
	}.Build()
	File_flyt_justeat_operational_alert_threshold_breached_proto = out.File
	file_flyt_justeat_operational_alert_threshold_breached_proto_rawDesc = nil
	file_flyt_justeat_operational_alert_threshold_breached_proto_goTypes = nil
	file_flyt_justeat_operational_alert_threshold_breached_proto_depIdxs = nil
}