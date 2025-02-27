// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: payments/getpaymentoptions.proto

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

type GetpaymentoptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32    `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ClientId   string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Amount     *float64 `protobuf:"fixed64,3,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
}

func (x *GetpaymentoptionsRequest) Reset() {
	*x = GetpaymentoptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_getpaymentoptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetpaymentoptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetpaymentoptionsRequest) ProtoMessage() {}

func (x *GetpaymentoptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_getpaymentoptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetpaymentoptionsRequest.ProtoReflect.Descriptor instead.
func (*GetpaymentoptionsRequest) Descriptor() ([]byte, []int) {
	return file_payments_getpaymentoptions_proto_rawDescGZIP(), []int{0}
}

func (x *GetpaymentoptionsRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetpaymentoptionsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetpaymentoptionsRequest) GetAmount() float64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

type GetpaymentoptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string                            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32                             `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ClientId   string                            `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CustomerId int32                             `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Data       []*GetpaymentoptionsResponse_Data `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
	Message    *string                           `protobuf:"bytes,6,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *GetpaymentoptionsResponse) Reset() {
	*x = GetpaymentoptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_getpaymentoptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetpaymentoptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetpaymentoptionsResponse) ProtoMessage() {}

func (x *GetpaymentoptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_getpaymentoptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetpaymentoptionsResponse.ProtoReflect.Descriptor instead.
func (*GetpaymentoptionsResponse) Descriptor() ([]byte, []int) {
	return file_payments_getpaymentoptions_proto_rawDescGZIP(), []int{1}
}

func (x *GetpaymentoptionsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetpaymentoptionsResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetpaymentoptionsResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetpaymentoptionsResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetpaymentoptionsResponse) GetData() []*GetpaymentoptionsResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetpaymentoptionsResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type GetpaymentoptionsResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway     string   `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Amount      float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ServiceFee  float64  `protobuf:"fixed64,3,opt,name=service_fee,json=serviceFee,proto3" json:"service_fee,omitempty"`
	OtherFees   float64  `protobuf:"fixed64,4,opt,name=other_fees,json=otherFees,proto3" json:"other_fees,omitempty"`
	TotalAmount float64  `protobuf:"fixed64,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Modes       []string `protobuf:"bytes,6,rep,name=modes,proto3" json:"modes,omitempty"`
}

func (x *GetpaymentoptionsResponse_Data) Reset() {
	*x = GetpaymentoptionsResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_getpaymentoptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetpaymentoptionsResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetpaymentoptionsResponse_Data) ProtoMessage() {}

func (x *GetpaymentoptionsResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_payments_getpaymentoptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetpaymentoptionsResponse_Data.ProtoReflect.Descriptor instead.
func (*GetpaymentoptionsResponse_Data) Descriptor() ([]byte, []int) {
	return file_payments_getpaymentoptions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetpaymentoptionsResponse_Data) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *GetpaymentoptionsResponse_Data) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetpaymentoptionsResponse_Data) GetServiceFee() float64 {
	if x != nil {
		return x.ServiceFee
	}
	return 0
}

func (x *GetpaymentoptionsResponse_Data) GetOtherFees() float64 {
	if x != nil {
		return x.OtherFees
	}
	return 0
}

func (x *GetpaymentoptionsResponse_Data) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *GetpaymentoptionsResponse_Data) GetModes() []string {
	if x != nil {
		return x.Modes
	}
	return nil
}

var File_payments_getpaymentoptions_proto protoreflect.FileDescriptor

var file_payments_getpaymentoptions_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x80,
	0x01, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xc1, 0x03, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x1a, 0xb1, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payments_getpaymentoptions_proto_rawDescOnce sync.Once
	file_payments_getpaymentoptions_proto_rawDescData = file_payments_getpaymentoptions_proto_rawDesc
)

func file_payments_getpaymentoptions_proto_rawDescGZIP() []byte {
	file_payments_getpaymentoptions_proto_rawDescOnce.Do(func() {
		file_payments_getpaymentoptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_getpaymentoptions_proto_rawDescData)
	})
	return file_payments_getpaymentoptions_proto_rawDescData
}

var file_payments_getpaymentoptions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payments_getpaymentoptions_proto_goTypes = []interface{}{
	(*GetpaymentoptionsRequest)(nil),       // 0: payments.getpaymentoptions.getpaymentoptionsRequest
	(*GetpaymentoptionsResponse)(nil),      // 1: payments.getpaymentoptions.getpaymentoptionsResponse
	(*GetpaymentoptionsResponse_Data)(nil), // 2: payments.getpaymentoptions.getpaymentoptionsResponse.Data
}
var file_payments_getpaymentoptions_proto_depIdxs = []int32{
	2, // 0: payments.getpaymentoptions.getpaymentoptionsResponse.data:type_name -> payments.getpaymentoptions.getpaymentoptionsResponse.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payments_getpaymentoptions_proto_init() }
func file_payments_getpaymentoptions_proto_init() {
	if File_payments_getpaymentoptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_getpaymentoptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetpaymentoptionsRequest); i {
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
		file_payments_getpaymentoptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetpaymentoptionsResponse); i {
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
		file_payments_getpaymentoptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetpaymentoptionsResponse_Data); i {
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
	file_payments_getpaymentoptions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_payments_getpaymentoptions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payments_getpaymentoptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_getpaymentoptions_proto_goTypes,
		DependencyIndexes: file_payments_getpaymentoptions_proto_depIdxs,
		MessageInfos:      file_payments_getpaymentoptions_proto_msgTypes,
	}.Build()
	File_payments_getpaymentoptions_proto = out.File
	file_payments_getpaymentoptions_proto_rawDesc = nil
	file_payments_getpaymentoptions_proto_goTypes = nil
	file_payments_getpaymentoptions_proto_depIdxs = nil
}
