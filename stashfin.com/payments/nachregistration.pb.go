// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: payments/nachregistration.proto

package payments

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

type MandatePartner int32

const (
	MandatePartner_INVALID  MandatePartner = 0
	MandatePartner_LOTUSPAY MandatePartner = 1
	MandatePartner_DIGIO    MandatePartner = 2
)

// Enum value maps for MandatePartner.
var (
	MandatePartner_name = map[int32]string{
		0: "INVALID",
		1: "LOTUSPAY",
		2: "DIGIO",
	}
	MandatePartner_value = map[string]int32{
		"INVALID":  0,
		"LOTUSPAY": 1,
		"DIGIO":    2,
	}
)

func (x MandatePartner) Enum() *MandatePartner {
	p := new(MandatePartner)
	*p = x
	return p
}

func (x MandatePartner) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MandatePartner) Descriptor() protoreflect.EnumDescriptor {
	return file_payments_nachregistration_proto_enumTypes[0].Descriptor()
}

func (MandatePartner) Type() protoreflect.EnumType {
	return &file_payments_nachregistration_proto_enumTypes[0]
}

func (x MandatePartner) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MandatePartner.Descriptor instead.
func (MandatePartner) EnumDescriptor() ([]byte, []int) {
	return file_payments_nachregistration_proto_rawDescGZIP(), []int{0}
}

type NachRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId        int64           `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ClientId          string          `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	MandateType       string          `protobuf:"bytes,3,opt,name=mandate_type,json=mandateType,proto3" json:"mandate_type,omitempty"`
	BankName          *string         `protobuf:"bytes,4,opt,name=bank_name,json=bankName,proto3,oneof" json:"bank_name,omitempty"`
	BankAccountNumber *string         `protobuf:"bytes,5,opt,name=bank_account_number,json=bankAccountNumber,proto3,oneof" json:"bank_account_number,omitempty"`
	BankIfsc          *string         `protobuf:"bytes,6,opt,name=bank_ifsc,json=bankIfsc,proto3,oneof" json:"bank_ifsc,omitempty"`
	MaximumAmount     float32         `protobuf:"fixed32,7,opt,name=maximum_amount,json=maximumAmount,proto3" json:"maximum_amount,omitempty"`
	ExternalOrderId   string          `protobuf:"bytes,8,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	VpaId             *string         `protobuf:"bytes,9,opt,name=vpa_id,json=vpaId,proto3,oneof" json:"vpa_id,omitempty"`
	GatewayIdentifier *MandatePartner `protobuf:"varint,10,opt,name=gatewayIdentifier,proto3,enum=payments.nachregistration.MandatePartner,oneof" json:"gatewayIdentifier,omitempty"`
}

func (x *NachRegistrationRequest) Reset() {
	*x = NachRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_nachregistration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NachRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NachRegistrationRequest) ProtoMessage() {}

func (x *NachRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_nachregistration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NachRegistrationRequest.ProtoReflect.Descriptor instead.
func (*NachRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_payments_nachregistration_proto_rawDescGZIP(), []int{0}
}

func (x *NachRegistrationRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *NachRegistrationRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *NachRegistrationRequest) GetMandateType() string {
	if x != nil {
		return x.MandateType
	}
	return ""
}

func (x *NachRegistrationRequest) GetBankName() string {
	if x != nil && x.BankName != nil {
		return *x.BankName
	}
	return ""
}

func (x *NachRegistrationRequest) GetBankAccountNumber() string {
	if x != nil && x.BankAccountNumber != nil {
		return *x.BankAccountNumber
	}
	return ""
}

func (x *NachRegistrationRequest) GetBankIfsc() string {
	if x != nil && x.BankIfsc != nil {
		return *x.BankIfsc
	}
	return ""
}

func (x *NachRegistrationRequest) GetMaximumAmount() float32 {
	if x != nil {
		return x.MaximumAmount
	}
	return 0
}

func (x *NachRegistrationRequest) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

func (x *NachRegistrationRequest) GetVpaId() string {
	if x != nil && x.VpaId != nil {
		return *x.VpaId
	}
	return ""
}

func (x *NachRegistrationRequest) GetGatewayIdentifier() MandatePartner {
	if x != nil && x.GatewayIdentifier != nil {
		return *x.GatewayIdentifier
	}
	return MandatePartner_INVALID
}

type NachRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string                         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32                          `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Data       *NachRegistrationResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NachRegistrationResponse) Reset() {
	*x = NachRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_nachregistration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NachRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NachRegistrationResponse) ProtoMessage() {}

func (x *NachRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_nachregistration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NachRegistrationResponse.ProtoReflect.Descriptor instead.
func (*NachRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_payments_nachregistration_proto_rawDescGZIP(), []int{1}
}

func (x *NachRegistrationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NachRegistrationResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *NachRegistrationResponse) GetData() *NachRegistrationResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type NachRegistrationResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *NachRegistrationResponse_Data) Reset() {
	*x = NachRegistrationResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_nachregistration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NachRegistrationResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NachRegistrationResponse_Data) ProtoMessage() {}

func (x *NachRegistrationResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_payments_nachregistration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NachRegistrationResponse_Data.ProtoReflect.Descriptor instead.
func (*NachRegistrationResponse_Data) Descriptor() ([]byte, []int) {
	return file_payments_nachregistration_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NachRegistrationResponse_Data) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_payments_nachregistration_proto protoreflect.FileDescriptor

var file_payments_nachregistration_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6e, 0x61, 0x63, 0x68, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x95, 0x04, 0x0a,
	0x17, 0x6e, 0x61, 0x63, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61,
	0x6e, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x11, 0x62, 0x61, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x66, 0x73, 0x63, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x66, 0x73, 0x63, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x76, 0x70, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x76, 0x70, 0x61, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x5c, 0x0a, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x48, 0x04, 0x52, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x16, 0x0a,
	0x14, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69,
	0x66, 0x73, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x76, 0x70, 0x61, 0x5f, 0x69, 0x64, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0xcc, 0x01, 0x0a, 0x18, 0x6e, 0x61, 0x63, 0x68, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x29, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x2a, 0x36, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x54, 0x55, 0x53, 0x50, 0x41, 0x59, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x49, 0x47, 0x49, 0x4f, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_payments_nachregistration_proto_rawDescOnce sync.Once
	file_payments_nachregistration_proto_rawDescData = file_payments_nachregistration_proto_rawDesc
)

func file_payments_nachregistration_proto_rawDescGZIP() []byte {
	file_payments_nachregistration_proto_rawDescOnce.Do(func() {
		file_payments_nachregistration_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_nachregistration_proto_rawDescData)
	})
	return file_payments_nachregistration_proto_rawDescData
}

var file_payments_nachregistration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payments_nachregistration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payments_nachregistration_proto_goTypes = []interface{}{
	(MandatePartner)(0),                   // 0: payments.nachregistration.MandatePartner
	(*NachRegistrationRequest)(nil),       // 1: payments.nachregistration.nachRegistrationRequest
	(*NachRegistrationResponse)(nil),      // 2: payments.nachregistration.nachRegistrationResponse
	(*NachRegistrationResponse_Data)(nil), // 3: payments.nachregistration.nachRegistrationResponse.Data
}
var file_payments_nachregistration_proto_depIdxs = []int32{
	0, // 0: payments.nachregistration.nachRegistrationRequest.gatewayIdentifier:type_name -> payments.nachregistration.MandatePartner
	3, // 1: payments.nachregistration.nachRegistrationResponse.data:type_name -> payments.nachregistration.nachRegistrationResponse.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payments_nachregistration_proto_init() }
func file_payments_nachregistration_proto_init() {
	if File_payments_nachregistration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_nachregistration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NachRegistrationRequest); i {
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
		file_payments_nachregistration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NachRegistrationResponse); i {
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
		file_payments_nachregistration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NachRegistrationResponse_Data); i {
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
	file_payments_nachregistration_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payments_nachregistration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_nachregistration_proto_goTypes,
		DependencyIndexes: file_payments_nachregistration_proto_depIdxs,
		EnumInfos:         file_payments_nachregistration_proto_enumTypes,
		MessageInfos:      file_payments_nachregistration_proto_msgTypes,
	}.Build()
	File_payments_nachregistration_proto = out.File
	file_payments_nachregistration_proto_rawDesc = nil
	file_payments_nachregistration_proto_goTypes = nil
	file_payments_nachregistration_proto_depIdxs = nil
}
