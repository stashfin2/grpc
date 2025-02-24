// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: payments/initiatepayment.proto

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

type InitiatepaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      int64   `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ClientId        string  `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Mode            string  `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Amount          int64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ExternalOrderId string  `protobuf:"bytes,5,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	MetaData        *string `protobuf:"bytes,6,opt,name=meta_data,json=metaData,proto3,oneof" json:"meta_data,omitempty"`
}

func (x *InitiatepaymentRequest) Reset() {
	*x = InitiatepaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_initiatepayment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatepaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatepaymentRequest) ProtoMessage() {}

func (x *InitiatepaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_initiatepayment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatepaymentRequest.ProtoReflect.Descriptor instead.
func (*InitiatepaymentRequest) Descriptor() ([]byte, []int) {
	return file_payments_initiatepayment_proto_rawDescGZIP(), []int{0}
}

func (x *InitiatepaymentRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *InitiatepaymentRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *InitiatepaymentRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *InitiatepaymentRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InitiatepaymentRequest) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

func (x *InitiatepaymentRequest) GetMetaData() string {
	if x != nil && x.MetaData != nil {
		return *x.MetaData
	}
	return ""
}

type InitiatepaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string                        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32                         `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Data       *InitiatepaymentResponse_Data `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *InitiatepaymentResponse) Reset() {
	*x = InitiatepaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_initiatepayment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatepaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatepaymentResponse) ProtoMessage() {}

func (x *InitiatepaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_initiatepayment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatepaymentResponse.ProtoReflect.Descriptor instead.
func (*InitiatepaymentResponse) Descriptor() ([]byte, []int) {
	return file_payments_initiatepayment_proto_rawDescGZIP(), []int{1}
}

func (x *InitiatepaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InitiatepaymentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *InitiatepaymentResponse) GetData() *InitiatepaymentResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type InitiatepaymentResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl *string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3,oneof" json:"redirect_url,omitempty"`
}

func (x *InitiatepaymentResponse_Data) Reset() {
	*x = InitiatepaymentResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_initiatepayment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatepaymentResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatepaymentResponse_Data) ProtoMessage() {}

func (x *InitiatepaymentResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_payments_initiatepayment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatepaymentResponse_Data.ProtoReflect.Descriptor instead.
func (*InitiatepaymentResponse_Data) Descriptor() ([]byte, []int) {
	return file_payments_initiatepayment_proto_rawDescGZIP(), []int{1, 0}
}

func (x *InitiatepaymentResponse_Data) GetRedirectUrl() string {
	if x != nil && x.RedirectUrl != nil {
		return *x.RedirectUrl
	}
	return ""
}

var File_payments_initiatepayment_proto protoreflect.FileDescriptor

var file_payments_initiatepayment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xed, 0x01, 0x0a, 0x17,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x4f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x1a, 0x3f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x88, 0x01,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_payments_initiatepayment_proto_rawDescOnce sync.Once
	file_payments_initiatepayment_proto_rawDescData = file_payments_initiatepayment_proto_rawDesc
)

func file_payments_initiatepayment_proto_rawDescGZIP() []byte {
	file_payments_initiatepayment_proto_rawDescOnce.Do(func() {
		file_payments_initiatepayment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_initiatepayment_proto_rawDescData)
	})
	return file_payments_initiatepayment_proto_rawDescData
}

var file_payments_initiatepayment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payments_initiatepayment_proto_goTypes = []interface{}{
	(*InitiatepaymentRequest)(nil),       // 0: payments.initiatepayment.initiatepaymentRequest
	(*InitiatepaymentResponse)(nil),      // 1: payments.initiatepayment.initiatepaymentResponse
	(*InitiatepaymentResponse_Data)(nil), // 2: payments.initiatepayment.initiatepaymentResponse.Data
}
var file_payments_initiatepayment_proto_depIdxs = []int32{
	2, // 0: payments.initiatepayment.initiatepaymentResponse.data:type_name -> payments.initiatepayment.initiatepaymentResponse.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payments_initiatepayment_proto_init() }
func file_payments_initiatepayment_proto_init() {
	if File_payments_initiatepayment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_initiatepayment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatepaymentRequest); i {
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
		file_payments_initiatepayment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatepaymentResponse); i {
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
		file_payments_initiatepayment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatepaymentResponse_Data); i {
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
	file_payments_initiatepayment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_payments_initiatepayment_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_payments_initiatepayment_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payments_initiatepayment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_initiatepayment_proto_goTypes,
		DependencyIndexes: file_payments_initiatepayment_proto_depIdxs,
		MessageInfos:      file_payments_initiatepayment_proto_msgTypes,
	}.Build()
	File_payments_initiatepayment_proto = out.File
	file_payments_initiatepayment_proto_rawDesc = nil
	file_payments_initiatepayment_proto_goTypes = nil
	file_payments_initiatepayment_proto_depIdxs = nil
}
