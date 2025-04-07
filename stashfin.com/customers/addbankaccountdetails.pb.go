// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/addbankaccountdetails.proto

package customers

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

type AddBankAccountDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	BankName      string `protobuf:"bytes,2,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	IfscCode      string `protobuf:"bytes,3,opt,name=ifsc_code,json=ifscCode,proto3" json:"ifsc_code,omitempty"`
}

func (x *AddBankAccountDetailsRequest) Reset() {
	*x = AddBankAccountDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_addbankaccountdetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBankAccountDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBankAccountDetailsRequest) ProtoMessage() {}

func (x *AddBankAccountDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_addbankaccountdetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBankAccountDetailsRequest.ProtoReflect.Descriptor instead.
func (*AddBankAccountDetailsRequest) Descriptor() ([]byte, []int) {
	return file_customers_addbankaccountdetails_proto_rawDescGZIP(), []int{0}
}

func (x *AddBankAccountDetailsRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *AddBankAccountDetailsRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *AddBankAccountDetailsRequest) GetIfscCode() string {
	if x != nil {
		return x.IfscCode
	}
	return ""
}

type AddBankAccountDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddBankAccountDetailsResponse) Reset() {
	*x = AddBankAccountDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_addbankaccountdetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBankAccountDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBankAccountDetailsResponse) ProtoMessage() {}

func (x *AddBankAccountDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_addbankaccountdetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBankAccountDetailsResponse.ProtoReflect.Descriptor instead.
func (*AddBankAccountDetailsResponse) Descriptor() ([]byte, []int) {
	return file_customers_addbankaccountdetails_proto_rawDescGZIP(), []int{1}
}

func (x *AddBankAccountDetailsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_customers_addbankaccountdetails_proto protoreflect.FileDescriptor

var file_customers_addbankaccountdetails_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x62,
	0x61, 0x6e, 0x6b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x62, 0x61, 0x6e, 0x6b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x7f, 0x0a, 0x1c, 0x61, 0x64, 0x64, 0x42,
	0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x66, 0x73, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x66, 0x73, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2f, 0x0a, 0x1d, 0x61, 0x64, 0x64,
	0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_customers_addbankaccountdetails_proto_rawDescOnce sync.Once
	file_customers_addbankaccountdetails_proto_rawDescData = file_customers_addbankaccountdetails_proto_rawDesc
)

func file_customers_addbankaccountdetails_proto_rawDescGZIP() []byte {
	file_customers_addbankaccountdetails_proto_rawDescOnce.Do(func() {
		file_customers_addbankaccountdetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_addbankaccountdetails_proto_rawDescData)
	})
	return file_customers_addbankaccountdetails_proto_rawDescData
}

var file_customers_addbankaccountdetails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_addbankaccountdetails_proto_goTypes = []interface{}{
	(*AddBankAccountDetailsRequest)(nil),  // 0: customers.addbankaccountdetails.addBankAccountDetailsRequest
	(*AddBankAccountDetailsResponse)(nil), // 1: customers.addbankaccountdetails.addBankAccountDetailsResponse
}
var file_customers_addbankaccountdetails_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_addbankaccountdetails_proto_init() }
func file_customers_addbankaccountdetails_proto_init() {
	if File_customers_addbankaccountdetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_addbankaccountdetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBankAccountDetailsRequest); i {
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
		file_customers_addbankaccountdetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBankAccountDetailsResponse); i {
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
			RawDescriptor: file_customers_addbankaccountdetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_addbankaccountdetails_proto_goTypes,
		DependencyIndexes: file_customers_addbankaccountdetails_proto_depIdxs,
		MessageInfos:      file_customers_addbankaccountdetails_proto_msgTypes,
	}.Build()
	File_customers_addbankaccountdetails_proto = out.File
	file_customers_addbankaccountdetails_proto_rawDesc = nil
	file_customers_addbankaccountdetails_proto_goTypes = nil
	file_customers_addbankaccountdetails_proto_depIdxs = nil
}
