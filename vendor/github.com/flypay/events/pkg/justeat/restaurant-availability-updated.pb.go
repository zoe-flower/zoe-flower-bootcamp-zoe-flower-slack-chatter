// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: justeat/restaurant-availability-updated.proto

package justeat

import (
	_ "github.com/flypay/events/pkg/flyt"
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

type RestaurantAvailabilityUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId                          string            `protobuf:"bytes,1,opt,name=restaurant_id,json=RestaurantId,proto3" json:"restaurant_id,omitempty"`
	IsTempOffline                         bool              `protobuf:"varint,2,opt,name=is_temp_offline,json=IsTempOffline,proto3" json:"is_temp_offline,omitempty"`
	DisableCollection                     bool              `protobuf:"varint,3,opt,name=disable_collection,json=DisableCollection,proto3" json:"disable_collection,omitempty"`
	PreOrderForFutureAvailabilityDisabled bool              `protobuf:"varint,4,opt,name=pre_order_for_future_availability_disabled,json=PreOrderForFutureAvailabilityDisabled,proto3" json:"pre_order_for_future_availability_disabled,omitempty"`
	PreOrderAtCheckoutDisabled            bool              `protobuf:"varint,5,opt,name=pre_order_at_checkout_disabled,json=PreOrderAtCheckoutDisabled,proto3" json:"pre_order_at_checkout_disabled,omitempty"`
	DeliveryBusyness                      *DeliveryBusyness `protobuf:"bytes,6,opt,name=delivery_busyness,json=DeliveryBusyness,proto3" json:"delivery_busyness,omitempty"`
	DateUpdated                           string            `protobuf:"bytes,7,opt,name=date_updated,json=DateUpdated,proto3" json:"date_updated,omitempty"`
	Tenant                                string            `protobuf:"bytes,8,opt,name=tenant,json=Tenant,proto3" json:"tenant,omitempty"`
}

func (x *RestaurantAvailabilityUpdated) Reset() {
	*x = RestaurantAvailabilityUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantAvailabilityUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantAvailabilityUpdated) ProtoMessage() {}

func (x *RestaurantAvailabilityUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantAvailabilityUpdated.ProtoReflect.Descriptor instead.
func (*RestaurantAvailabilityUpdated) Descriptor() ([]byte, []int) {
	return file_justeat_restaurant_availability_updated_proto_rawDescGZIP(), []int{0}
}

func (x *RestaurantAvailabilityUpdated) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *RestaurantAvailabilityUpdated) GetIsTempOffline() bool {
	if x != nil {
		return x.IsTempOffline
	}
	return false
}

func (x *RestaurantAvailabilityUpdated) GetDisableCollection() bool {
	if x != nil {
		return x.DisableCollection
	}
	return false
}

func (x *RestaurantAvailabilityUpdated) GetPreOrderForFutureAvailabilityDisabled() bool {
	if x != nil {
		return x.PreOrderForFutureAvailabilityDisabled
	}
	return false
}

func (x *RestaurantAvailabilityUpdated) GetPreOrderAtCheckoutDisabled() bool {
	if x != nil {
		return x.PreOrderAtCheckoutDisabled
	}
	return false
}

func (x *RestaurantAvailabilityUpdated) GetDeliveryBusyness() *DeliveryBusyness {
	if x != nil {
		return x.DeliveryBusyness
	}
	return nil
}

func (x *RestaurantAvailabilityUpdated) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

func (x *RestaurantAvailabilityUpdated) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

type DeliveryBusyness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryTimeMinutes *DeliveryBusynessDeliveryTime `protobuf:"bytes,1,opt,name=delivery_time_minutes,json=DeliveryTimeMinutes,proto3" json:"delivery_time_minutes,omitempty"`
	Metadata            *DeliveryBusynessMetadata     `protobuf:"bytes,2,opt,name=metadata,json=Metadata,proto3" json:"metadata,omitempty"`
	State               string                        `protobuf:"bytes,3,opt,name=state,json=State,proto3" json:"state,omitempty"`
}

func (x *DeliveryBusyness) Reset() {
	*x = DeliveryBusyness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryBusyness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryBusyness) ProtoMessage() {}

func (x *DeliveryBusyness) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryBusyness.ProtoReflect.Descriptor instead.
func (*DeliveryBusyness) Descriptor() ([]byte, []int) {
	return file_justeat_restaurant_availability_updated_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryBusyness) GetDeliveryTimeMinutes() *DeliveryBusynessDeliveryTime {
	if x != nil {
		return x.DeliveryTimeMinutes
	}
	return nil
}

func (x *DeliveryBusyness) GetMetadata() *DeliveryBusynessMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DeliveryBusyness) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type DeliveryBusynessDeliveryTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestGuess   int32 `protobuf:"varint,1,opt,name=best_guess,json=BestGuess,proto3" json:"best_guess,omitempty"`
	Optimistic  int32 `protobuf:"varint,2,opt,name=optimistic,json=Optimistic,proto3" json:"optimistic,omitempty"`
	Pessimistic int32 `protobuf:"varint,3,opt,name=pessimistic,json=Pessimistic,proto3" json:"pessimistic,omitempty"`
	LowerBound  int32 `protobuf:"varint,4,opt,name=lower_bound,json=LowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound  int32 `protobuf:"varint,5,opt,name=upper_bound,json=UpperBound,proto3" json:"upper_bound,omitempty"`
}

func (x *DeliveryBusynessDeliveryTime) Reset() {
	*x = DeliveryBusynessDeliveryTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryBusynessDeliveryTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryBusynessDeliveryTime) ProtoMessage() {}

func (x *DeliveryBusynessDeliveryTime) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryBusynessDeliveryTime.ProtoReflect.Descriptor instead.
func (*DeliveryBusynessDeliveryTime) Descriptor() ([]byte, []int) {
	return file_justeat_restaurant_availability_updated_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryBusynessDeliveryTime) GetBestGuess() int32 {
	if x != nil {
		return x.BestGuess
	}
	return 0
}

func (x *DeliveryBusynessDeliveryTime) GetOptimistic() int32 {
	if x != nil {
		return x.Optimistic
	}
	return 0
}

func (x *DeliveryBusynessDeliveryTime) GetPessimistic() int32 {
	if x != nil {
		return x.Pessimistic
	}
	return 0
}

func (x *DeliveryBusynessDeliveryTime) GetLowerBound() int32 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *DeliveryBusynessDeliveryTime) GetUpperBound() int32 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

type DeliveryBusynessMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason        string `protobuf:"bytes,1,opt,name=reason,json=Reason,proto3" json:"reason,omitempty"`
	CorrelationId string `protobuf:"bytes,2,opt,name=correlation_id,json=CorrelationId,proto3" json:"correlation_id,omitempty"`
}

func (x *DeliveryBusynessMetadata) Reset() {
	*x = DeliveryBusynessMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryBusynessMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryBusynessMetadata) ProtoMessage() {}

func (x *DeliveryBusynessMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_restaurant_availability_updated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryBusynessMetadata.ProtoReflect.Descriptor instead.
func (*DeliveryBusynessMetadata) Descriptor() ([]byte, []int) {
	return file_justeat_restaurant_availability_updated_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryBusynessMetadata) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *DeliveryBusynessMetadata) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

var File_justeat_restaurant_availability_updated_proto protoreflect.FileDescriptor

var file_justeat_restaurant_availability_updated_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x04, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x2a, 0x70, 0x72, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x25, 0x50, 0x72, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x42, 0x0a, 0x1e, 0x70, 0x72, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x50, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x41, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x62, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x53, 0x82, 0xb5, 0x18, 0x27, 0x6a, 0x75, 0x73, 0x74,
	0x65, 0x61, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0xa2, 0xbb, 0x18, 0x1d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0xaa, 0xbb, 0x18, 0x03, 0x61, 0x6c, 0x6c, 0x22, 0xc2, 0x01, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x59, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xc1, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x75, 0x73,
	0x79, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x42, 0x65, 0x73, 0x74, 0x47, 0x75, 0x65, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x65, 0x73, 0x73, 0x69, 0x6d, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x55, 0x70, 0x70, 0x65, 0x72, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x59, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x42, 0x75, 0x73, 0x79, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x93, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x42,
	0x22, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x4a, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xca, 0x02, 0x07, 0x4a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0xe2, 0x02, 0x13, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_justeat_restaurant_availability_updated_proto_rawDescOnce sync.Once
	file_justeat_restaurant_availability_updated_proto_rawDescData = file_justeat_restaurant_availability_updated_proto_rawDesc
)

func file_justeat_restaurant_availability_updated_proto_rawDescGZIP() []byte {
	file_justeat_restaurant_availability_updated_proto_rawDescOnce.Do(func() {
		file_justeat_restaurant_availability_updated_proto_rawDescData = protoimpl.X.CompressGZIP(file_justeat_restaurant_availability_updated_proto_rawDescData)
	})
	return file_justeat_restaurant_availability_updated_proto_rawDescData
}

var file_justeat_restaurant_availability_updated_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_justeat_restaurant_availability_updated_proto_goTypes = []interface{}{
	(*RestaurantAvailabilityUpdated)(nil), // 0: justeat.RestaurantAvailabilityUpdated
	(*DeliveryBusyness)(nil),              // 1: justeat.DeliveryBusyness
	(*DeliveryBusynessDeliveryTime)(nil),  // 2: justeat.DeliveryBusynessDeliveryTime
	(*DeliveryBusynessMetadata)(nil),      // 3: justeat.DeliveryBusynessMetadata
}
var file_justeat_restaurant_availability_updated_proto_depIdxs = []int32{
	1, // 0: justeat.RestaurantAvailabilityUpdated.delivery_busyness:type_name -> justeat.DeliveryBusyness
	2, // 1: justeat.DeliveryBusyness.delivery_time_minutes:type_name -> justeat.DeliveryBusynessDeliveryTime
	3, // 2: justeat.DeliveryBusyness.metadata:type_name -> justeat.DeliveryBusynessMetadata
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_justeat_restaurant_availability_updated_proto_init() }
func file_justeat_restaurant_availability_updated_proto_init() {
	if File_justeat_restaurant_availability_updated_proto != nil {
		return
	}
	file_justeat_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_justeat_restaurant_availability_updated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantAvailabilityUpdated); i {
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
		file_justeat_restaurant_availability_updated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryBusyness); i {
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
		file_justeat_restaurant_availability_updated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryBusynessDeliveryTime); i {
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
		file_justeat_restaurant_availability_updated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryBusynessMetadata); i {
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
			RawDescriptor: file_justeat_restaurant_availability_updated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_justeat_restaurant_availability_updated_proto_goTypes,
		DependencyIndexes: file_justeat_restaurant_availability_updated_proto_depIdxs,
		MessageInfos:      file_justeat_restaurant_availability_updated_proto_msgTypes,
	}.Build()
	File_justeat_restaurant_availability_updated_proto = out.File
	file_justeat_restaurant_availability_updated_proto_rawDesc = nil
	file_justeat_restaurant_availability_updated_proto_goTypes = nil
	file_justeat_restaurant_availability_updated_proto_depIdxs = nil
}