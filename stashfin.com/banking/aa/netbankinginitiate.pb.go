// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: banking/aa/netbankinginitiate.proto

package aa

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

type InputPerfiosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *InputPerfiosRequest) Reset() {
	*x = InputPerfiosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_netbankinginitiate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputPerfiosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPerfiosRequest) ProtoMessage() {}

func (x *InputPerfiosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_netbankinginitiate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPerfiosRequest.ProtoReflect.Descriptor instead.
func (*InputPerfiosRequest) Descriptor() ([]byte, []int) {
	return file_banking_aa_netbankinginitiate_proto_rawDescGZIP(), []int{0}
}

func (x *InputPerfiosRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type OutputPerfiosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload       string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature     string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Url           string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	TransactionId string `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ReturnUrl     string `protobuf:"bytes,5,opt,name=return_url,json=returnUrl,proto3" json:"return_url,omitempty"`
}

func (x *OutputPerfiosResponse) Reset() {
	*x = OutputPerfiosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_netbankinginitiate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputPerfiosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputPerfiosResponse) ProtoMessage() {}

func (x *OutputPerfiosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_netbankinginitiate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputPerfiosResponse.ProtoReflect.Descriptor instead.
func (*OutputPerfiosResponse) Descriptor() ([]byte, []int) {
	return file_banking_aa_netbankinginitiate_proto_rawDescGZIP(), []int{1}
}

func (x *OutputPerfiosResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *OutputPerfiosResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *OutputPerfiosResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *OutputPerfiosResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *OutputPerfiosResponse) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

var File_banking_aa_netbankinginitiate_proto protoreflect.FileDescriptor

var file_banking_aa_netbankinginitiate_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x6e, 0x65, 0x74,
	0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x61, 0x2e, 0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x65, 0x72,
	0x66, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a,
	0x15, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x65, 0x72, 0x66, 0x69, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banking_aa_netbankinginitiate_proto_rawDescOnce sync.Once
	file_banking_aa_netbankinginitiate_proto_rawDescData = file_banking_aa_netbankinginitiate_proto_rawDesc
)

func file_banking_aa_netbankinginitiate_proto_rawDescGZIP() []byte {
	file_banking_aa_netbankinginitiate_proto_rawDescOnce.Do(func() {
		file_banking_aa_netbankinginitiate_proto_rawDescData = protoimpl.X.CompressGZIP(file_banking_aa_netbankinginitiate_proto_rawDescData)
	})
	return file_banking_aa_netbankinginitiate_proto_rawDescData
}

var file_banking_aa_netbankinginitiate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_banking_aa_netbankinginitiate_proto_goTypes = []interface{}{
	(*InputPerfiosRequest)(nil),   // 0: banking.aa.netbankinginitiate.inputPerfiosRequest
	(*OutputPerfiosResponse)(nil), // 1: banking.aa.netbankinginitiate.OutputPerfiosResponse
}
var file_banking_aa_netbankinginitiate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_banking_aa_netbankinginitiate_proto_init() }
func file_banking_aa_netbankinginitiate_proto_init() {
	if File_banking_aa_netbankinginitiate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banking_aa_netbankinginitiate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputPerfiosRequest); i {
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
		file_banking_aa_netbankinginitiate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputPerfiosResponse); i {
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
			RawDescriptor: file_banking_aa_netbankinginitiate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banking_aa_netbankinginitiate_proto_goTypes,
		DependencyIndexes: file_banking_aa_netbankinginitiate_proto_depIdxs,
		MessageInfos:      file_banking_aa_netbankinginitiate_proto_msgTypes,
	}.Build()
	File_banking_aa_netbankinginitiate_proto = out.File
	file_banking_aa_netbankinginitiate_proto_rawDesc = nil
	file_banking_aa_netbankinginitiate_proto_goTypes = nil
	file_banking_aa_netbankinginitiate_proto_depIdxs = nil
}
