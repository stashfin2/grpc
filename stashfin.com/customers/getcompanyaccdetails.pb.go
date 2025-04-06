// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/getcompanyaccdetails.proto

package customers

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

type GetCompanyAccDetailsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCompanyAccDetailsRequest) Reset() {
	*x = GetCompanyAccDetailsRequest{}
	mi := &file_customers_getcompanyaccdetails_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyAccDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyAccDetailsRequest) ProtoMessage() {}

func (x *GetCompanyAccDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getcompanyaccdetails_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyAccDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyAccDetailsRequest) Descriptor() ([]byte, []int) {
	return file_customers_getcompanyaccdetails_proto_rawDescGZIP(), []int{0}
}

type GetCompanyAccDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountNumber string                 `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	AccountName   string                 `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	IfscCode      string                 `protobuf:"bytes,3,opt,name=ifsc_code,json=ifscCode,proto3" json:"ifsc_code,omitempty"`
	BankName      string                 `protobuf:"bytes,4,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCompanyAccDetailsResponse) Reset() {
	*x = GetCompanyAccDetailsResponse{}
	mi := &file_customers_getcompanyaccdetails_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyAccDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyAccDetailsResponse) ProtoMessage() {}

func (x *GetCompanyAccDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getcompanyaccdetails_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyAccDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyAccDetailsResponse) Descriptor() ([]byte, []int) {
	return file_customers_getcompanyaccdetails_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyAccDetailsResponse) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *GetCompanyAccDetailsResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetCompanyAccDetailsResponse) GetIfscCode() string {
	if x != nil {
		return x.IfscCode
	}
	return ""
}

func (x *GetCompanyAccDetailsResponse) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

var File_customers_getcompanyaccdetails_proto protoreflect.FileDescriptor

var file_customers_getcompanyaccdetails_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x63, 0x63, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x63, 0x63, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x41, 0x63, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x1c, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x41, 0x63, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x66, 0x73, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x66, 0x73, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_customers_getcompanyaccdetails_proto_rawDescOnce sync.Once
	file_customers_getcompanyaccdetails_proto_rawDescData []byte
)

func file_customers_getcompanyaccdetails_proto_rawDescGZIP() []byte {
	file_customers_getcompanyaccdetails_proto_rawDescOnce.Do(func() {
		file_customers_getcompanyaccdetails_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_getcompanyaccdetails_proto_rawDesc), len(file_customers_getcompanyaccdetails_proto_rawDesc)))
	})
	return file_customers_getcompanyaccdetails_proto_rawDescData
}

var file_customers_getcompanyaccdetails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_getcompanyaccdetails_proto_goTypes = []any{
	(*GetCompanyAccDetailsRequest)(nil),  // 0: customers.getcompanyaccdetails.getCompanyAccDetailsRequest
	(*GetCompanyAccDetailsResponse)(nil), // 1: customers.getcompanyaccdetails.getCompanyAccDetailsResponse
}
var file_customers_getcompanyaccdetails_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_getcompanyaccdetails_proto_init() }
func file_customers_getcompanyaccdetails_proto_init() {
	if File_customers_getcompanyaccdetails_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_getcompanyaccdetails_proto_rawDesc), len(file_customers_getcompanyaccdetails_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_getcompanyaccdetails_proto_goTypes,
		DependencyIndexes: file_customers_getcompanyaccdetails_proto_depIdxs,
		MessageInfos:      file_customers_getcompanyaccdetails_proto_msgTypes,
	}.Build()
	File_customers_getcompanyaccdetails_proto = out.File
	file_customers_getcompanyaccdetails_proto_goTypes = nil
	file_customers_getcompanyaccdetails_proto_depIdxs = nil
}
