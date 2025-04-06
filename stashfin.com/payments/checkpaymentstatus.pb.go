// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: payments/checkpaymentstatus.proto

package payments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckPaymentStatusRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	CustomerId      int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ClientId        string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ExternalOrderId string                 `protobuf:"bytes,3,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CheckPaymentStatusRequest) Reset() {
	*x = CheckPaymentStatusRequest{}
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckPaymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPaymentStatusRequest) ProtoMessage() {}

func (x *CheckPaymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPaymentStatusRequest.ProtoReflect.Descriptor instead.
func (*CheckPaymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_payments_checkpaymentstatus_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPaymentStatusRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CheckPaymentStatusRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CheckPaymentStatusRequest) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

type CheckPaymentStatusResponse struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Status        string                           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode    int32                            `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Data          *CheckPaymentStatusResponse_Data `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckPaymentStatusResponse) Reset() {
	*x = CheckPaymentStatusResponse{}
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckPaymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPaymentStatusResponse) ProtoMessage() {}

func (x *CheckPaymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPaymentStatusResponse.ProtoReflect.Descriptor instead.
func (*CheckPaymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_payments_checkpaymentstatus_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPaymentStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CheckPaymentStatusResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CheckPaymentStatusResponse) GetData() *CheckPaymentStatusResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type CheckPaymentStatusResponse_Data struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	PaymentStatus   string                 `protobuf:"bytes,1,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	ExternalOrderId string                 `protobuf:"bytes,2,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CheckPaymentStatusResponse_Data) Reset() {
	*x = CheckPaymentStatusResponse_Data{}
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckPaymentStatusResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPaymentStatusResponse_Data) ProtoMessage() {}

func (x *CheckPaymentStatusResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_payments_checkpaymentstatus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPaymentStatusResponse_Data.ProtoReflect.Descriptor instead.
func (*CheckPaymentStatusResponse_Data) Descriptor() ([]byte, []int) {
	return file_payments_checkpaymentstatus_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CheckPaymentStatusResponse_Data) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *CheckPaymentStatusResponse_Data) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

var File_payments_checkpaymentstatus_proto protoreflect.FileDescriptor

var file_payments_checkpaymentstatus_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x85, 0x01, 0x0a, 0x19, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x90, 0x02, 0x0a, 0x1a, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x55, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x1a, 0x59, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_payments_checkpaymentstatus_proto_rawDescOnce sync.Once
	file_payments_checkpaymentstatus_proto_rawDescData []byte
)

func file_payments_checkpaymentstatus_proto_rawDescGZIP() []byte {
	file_payments_checkpaymentstatus_proto_rawDescOnce.Do(func() {
		file_payments_checkpaymentstatus_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payments_checkpaymentstatus_proto_rawDesc), len(file_payments_checkpaymentstatus_proto_rawDesc)))
	})
	return file_payments_checkpaymentstatus_proto_rawDescData
}

var file_payments_checkpaymentstatus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payments_checkpaymentstatus_proto_goTypes = []any{
	(*CheckPaymentStatusRequest)(nil),       // 0: payments.checkpaymentstatus.checkPaymentStatusRequest
	(*CheckPaymentStatusResponse)(nil),      // 1: payments.checkpaymentstatus.checkPaymentStatusResponse
	(*CheckPaymentStatusResponse_Data)(nil), // 2: payments.checkpaymentstatus.checkPaymentStatusResponse.Data
}
var file_payments_checkpaymentstatus_proto_depIdxs = []int32{
	2, // 0: payments.checkpaymentstatus.checkPaymentStatusResponse.data:type_name -> payments.checkpaymentstatus.checkPaymentStatusResponse.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payments_checkpaymentstatus_proto_init() }
func file_payments_checkpaymentstatus_proto_init() {
	if File_payments_checkpaymentstatus_proto != nil {
		return
	}
	file_payments_checkpaymentstatus_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payments_checkpaymentstatus_proto_rawDesc), len(file_payments_checkpaymentstatus_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_checkpaymentstatus_proto_goTypes,
		DependencyIndexes: file_payments_checkpaymentstatus_proto_depIdxs,
		MessageInfos:      file_payments_checkpaymentstatus_proto_msgTypes,
	}.Build()
	File_payments_checkpaymentstatus_proto = out.File
	file_payments_checkpaymentstatus_proto_goTypes = nil
	file_payments_checkpaymentstatus_proto_depIdxs = nil
}
