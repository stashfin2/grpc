// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: upi/getpendingcollectrequest.proto

package upi

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

type PendingCollectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile    string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	DeviceId  string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ProfileId string `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *PendingCollectRequest) Reset() {
	*x = PendingCollectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_getpendingcollectrequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingCollectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingCollectRequest) ProtoMessage() {}

func (x *PendingCollectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getpendingcollectrequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingCollectRequest.ProtoReflect.Descriptor instead.
func (*PendingCollectRequest) Descriptor() ([]byte, []int) {
	return file_upi_getpendingcollectrequest_proto_rawDescGZIP(), []int{0}
}

func (x *PendingCollectRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *PendingCollectRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *PendingCollectRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

type CollectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date                  *int64   `protobuf:"varint,1,opt,name=date,proto3,oneof" json:"date,omitempty"`
	Note                  *string  `protobuf:"bytes,2,opt,name=note,proto3,oneof" json:"note,omitempty"`
	Amount                *float32 `protobuf:"fixed32,3,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
	SeqNo                 *string  `protobuf:"bytes,4,opt,name=seqNo,proto3,oneof" json:"seqNo,omitempty"`
	PayeeMccType          *string  `protobuf:"bytes,5,opt,name=payeeMccType,proto3,oneof" json:"payeeMccType,omitempty"`
	ExpireAfter           *string  `protobuf:"bytes,6,opt,name=expireAfter,proto3,oneof" json:"expireAfter,omitempty"`
	PayeeVa               *string  `protobuf:"bytes,7,opt,name=payeeVa,proto3,oneof" json:"payeeVa,omitempty"`
	PayeeName             *string  `protobuf:"bytes,9,opt,name=payeeName,proto3,oneof" json:"payeeName,omitempty"`
	UpiTranlogId          *int64   `protobuf:"varint,10,opt,name=upiTranlogId,proto3,oneof" json:"upiTranlogId,omitempty"`
	PayerVa               *string  `protobuf:"bytes,11,opt,name=payerVa,proto3,oneof" json:"payerVa,omitempty"`
	RefUrl                *string  `protobuf:"bytes,12,opt,name=refUrl,proto3,oneof" json:"refUrl,omitempty"`
	PayeeMccCode          *string  `protobuf:"bytes,13,opt,name=payeeMccCode,proto3,oneof" json:"payeeMccCode,omitempty"`
	Direction             *string  `protobuf:"bytes,14,opt,name=direction,proto3,oneof" json:"direction,omitempty"`
	Status                *string  `protobuf:"bytes,15,opt,name=status,proto3,oneof" json:"status,omitempty"`
	PayeeVerifiedMerchant *string  `protobuf:"bytes,16,opt,name=payeeVerifiedMerchant,proto3,oneof" json:"payeeVerifiedMerchant,omitempty"`
	PayeeType             *string  `protobuf:"bytes,17,opt,name=payeeType,proto3,oneof" json:"payeeType,omitempty"`
}

func (x *CollectRequest) Reset() {
	*x = CollectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_getpendingcollectrequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectRequest) ProtoMessage() {}

func (x *CollectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getpendingcollectrequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectRequest.ProtoReflect.Descriptor instead.
func (*CollectRequest) Descriptor() ([]byte, []int) {
	return file_upi_getpendingcollectrequest_proto_rawDescGZIP(), []int{1}
}

func (x *CollectRequest) GetDate() int64 {
	if x != nil && x.Date != nil {
		return *x.Date
	}
	return 0
}

func (x *CollectRequest) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *CollectRequest) GetAmount() float32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *CollectRequest) GetSeqNo() string {
	if x != nil && x.SeqNo != nil {
		return *x.SeqNo
	}
	return ""
}

func (x *CollectRequest) GetPayeeMccType() string {
	if x != nil && x.PayeeMccType != nil {
		return *x.PayeeMccType
	}
	return ""
}

func (x *CollectRequest) GetExpireAfter() string {
	if x != nil && x.ExpireAfter != nil {
		return *x.ExpireAfter
	}
	return ""
}

func (x *CollectRequest) GetPayeeVa() string {
	if x != nil && x.PayeeVa != nil {
		return *x.PayeeVa
	}
	return ""
}

func (x *CollectRequest) GetPayeeName() string {
	if x != nil && x.PayeeName != nil {
		return *x.PayeeName
	}
	return ""
}

func (x *CollectRequest) GetUpiTranlogId() int64 {
	if x != nil && x.UpiTranlogId != nil {
		return *x.UpiTranlogId
	}
	return 0
}

func (x *CollectRequest) GetPayerVa() string {
	if x != nil && x.PayerVa != nil {
		return *x.PayerVa
	}
	return ""
}

func (x *CollectRequest) GetRefUrl() string {
	if x != nil && x.RefUrl != nil {
		return *x.RefUrl
	}
	return ""
}

func (x *CollectRequest) GetPayeeMccCode() string {
	if x != nil && x.PayeeMccCode != nil {
		return *x.PayeeMccCode
	}
	return ""
}

func (x *CollectRequest) GetDirection() string {
	if x != nil && x.Direction != nil {
		return *x.Direction
	}
	return ""
}

func (x *CollectRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *CollectRequest) GetPayeeVerifiedMerchant() string {
	if x != nil && x.PayeeVerifiedMerchant != nil {
		return *x.PayeeVerifiedMerchant
	}
	return ""
}

func (x *CollectRequest) GetPayeeType() string {
	if x != nil && x.PayeeType != nil {
		return *x.PayeeType
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*CollectRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_getpendingcollectrequest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getpendingcollectrequest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_upi_getpendingcollectrequest_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetRequests() []*CollectRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type PendingCollectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Data  `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *PendingCollectResponse) Reset() {
	*x = PendingCollectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_getpendingcollectrequest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingCollectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingCollectResponse) ProtoMessage() {}

func (x *PendingCollectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getpendingcollectrequest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingCollectResponse.ProtoReflect.Descriptor instead.
func (*PendingCollectResponse) Descriptor() ([]byte, []int) {
	return file_upi_getpendingcollectrequest_proto_rawDescGZIP(), []int{3}
}

func (x *PendingCollectResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PendingCollectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PendingCollectResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_getpendingcollectrequest_proto protoreflect.FileDescriptor

var file_upi_getpendingcollectrequest_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x6b, 0x0a, 0x15, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x94, 0x06, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x4d, 0x63, 0x63, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4d, 0x63, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x61, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0c, 0x75, 0x70, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x6c,
	0x6f, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x56, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x65,
	0x72, 0x56, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x55, 0x72, 0x6c,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x06, 0x72, 0x65, 0x66, 0x55, 0x72, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4d, 0x63, 0x63, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0c, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x4d, 0x63, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x0c, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x15,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0e, 0x52, 0x15, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0f, 0x52, 0x09, 0x70, 0x61,
	0x79, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x65, 0x71, 0x4e,
	0x6f, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4d, 0x63, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x61, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x75, 0x70, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x56, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65,
	0x66, 0x55, 0x72, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4d, 0x63,
	0x63, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x18,
	0x0a, 0x16, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x48,
	0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_upi_getpendingcollectrequest_proto_rawDescOnce sync.Once
	file_upi_getpendingcollectrequest_proto_rawDescData = file_upi_getpendingcollectrequest_proto_rawDesc
)

func file_upi_getpendingcollectrequest_proto_rawDescGZIP() []byte {
	file_upi_getpendingcollectrequest_proto_rawDescOnce.Do(func() {
		file_upi_getpendingcollectrequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_upi_getpendingcollectrequest_proto_rawDescData)
	})
	return file_upi_getpendingcollectrequest_proto_rawDescData
}

var file_upi_getpendingcollectrequest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upi_getpendingcollectrequest_proto_goTypes = []interface{}{
	(*PendingCollectRequest)(nil),  // 0: upi.getPendingCollectRequest.pendingCollectRequest
	(*CollectRequest)(nil),         // 1: upi.getPendingCollectRequest.CollectRequest
	(*Data)(nil),                   // 2: upi.getPendingCollectRequest.Data
	(*PendingCollectResponse)(nil), // 3: upi.getPendingCollectRequest.pendingCollectResponse
}
var file_upi_getpendingcollectrequest_proto_depIdxs = []int32{
	1, // 0: upi.getPendingCollectRequest.Data.requests:type_name -> upi.getPendingCollectRequest.CollectRequest
	2, // 1: upi.getPendingCollectRequest.pendingCollectResponse.data:type_name -> upi.getPendingCollectRequest.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upi_getpendingcollectrequest_proto_init() }
func file_upi_getpendingcollectrequest_proto_init() {
	if File_upi_getpendingcollectrequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upi_getpendingcollectrequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingCollectRequest); i {
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
		file_upi_getpendingcollectrequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectRequest); i {
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
		file_upi_getpendingcollectrequest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_upi_getpendingcollectrequest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingCollectResponse); i {
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
	file_upi_getpendingcollectrequest_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_upi_getpendingcollectrequest_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upi_getpendingcollectrequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_getpendingcollectrequest_proto_goTypes,
		DependencyIndexes: file_upi_getpendingcollectrequest_proto_depIdxs,
		MessageInfos:      file_upi_getpendingcollectrequest_proto_msgTypes,
	}.Build()
	File_upi_getpendingcollectrequest_proto = out.File
	file_upi_getpendingcollectrequest_proto_rawDesc = nil
	file_upi_getpendingcollectrequest_proto_goTypes = nil
	file_upi_getpendingcollectrequest_proto_depIdxs = nil
}
