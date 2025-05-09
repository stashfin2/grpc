// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: banking/aa/getcustomerbank.proto

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

type GetCustomerBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCustomerBankRequest) Reset() {
	*x = GetCustomerBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_getcustomerbank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerBankRequest) ProtoMessage() {}

func (x *GetCustomerBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_getcustomerbank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerBankRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerBankRequest) Descriptor() ([]byte, []int) {
	return file_banking_aa_getcustomerbank_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerBankRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetCustomerBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []*BankDetails `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCustomerBankResponse) Reset() {
	*x = GetCustomerBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_getcustomerbank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerBankResponse) ProtoMessage() {}

func (x *GetCustomerBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_getcustomerbank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerBankResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerBankResponse) Descriptor() ([]byte, []int) {
	return file_banking_aa_getcustomerbank_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerBankResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetCustomerBankResponse) GetData() []*BankDetails {
	if x != nil {
		return x.Data
	}
	return nil
}

type BankDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber *string `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3,oneof" json:"account_number,omitempty"`
	IfscCode      *string `protobuf:"bytes,2,opt,name=ifsc_code,json=ifscCode,proto3,oneof" json:"ifsc_code,omitempty"`
	CustomerName  *string `protobuf:"bytes,3,opt,name=customer_name,json=customerName,proto3,oneof" json:"customer_name,omitempty"`
	IsPrimary     *bool   `protobuf:"varint,4,opt,name=is_primary,json=isPrimary,proto3,oneof" json:"is_primary,omitempty"`
	BankLogoUrl   *string `protobuf:"bytes,5,opt,name=bank_logo_url,json=bankLogoUrl,proto3,oneof" json:"bank_logo_url,omitempty"`
}

func (x *BankDetails) Reset() {
	*x = BankDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_getcustomerbank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankDetails) ProtoMessage() {}

func (x *BankDetails) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_getcustomerbank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankDetails.ProtoReflect.Descriptor instead.
func (*BankDetails) Descriptor() ([]byte, []int) {
	return file_banking_aa_getcustomerbank_proto_rawDescGZIP(), []int{2}
}

func (x *BankDetails) GetAccountNumber() string {
	if x != nil && x.AccountNumber != nil {
		return *x.AccountNumber
	}
	return ""
}

func (x *BankDetails) GetIfscCode() string {
	if x != nil && x.IfscCode != nil {
		return *x.IfscCode
	}
	return ""
}

func (x *BankDetails) GetCustomerName() string {
	if x != nil && x.CustomerName != nil {
		return *x.CustomerName
	}
	return ""
}

func (x *BankDetails) GetIsPrimary() bool {
	if x != nil && x.IsPrimary != nil {
		return *x.IsPrimary
	}
	return false
}

func (x *BankDetails) GetBankLogoUrl() string {
	if x != nil && x.BankLogoUrl != nil {
		return *x.BankLogoUrl
	}
	return ""
}

var File_banking_aa_getcustomerbank_proto protoreflect.FileDescriptor

var file_banking_aa_getcustomerbank_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65, 0x74,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67,
	0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0x39,
	0x0a, 0x16, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x17, 0x67, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa6, 0x02, 0x0a, 0x0b, 0x42, 0x61,
	0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x66, 0x73, 0x63, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x69, 0x66, 0x73, 0x63,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x6f,
	0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0b,
	0x62, 0x61, 0x6e, 0x6b, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x66, 0x73, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banking_aa_getcustomerbank_proto_rawDescOnce sync.Once
	file_banking_aa_getcustomerbank_proto_rawDescData = file_banking_aa_getcustomerbank_proto_rawDesc
)

func file_banking_aa_getcustomerbank_proto_rawDescGZIP() []byte {
	file_banking_aa_getcustomerbank_proto_rawDescOnce.Do(func() {
		file_banking_aa_getcustomerbank_proto_rawDescData = protoimpl.X.CompressGZIP(file_banking_aa_getcustomerbank_proto_rawDescData)
	})
	return file_banking_aa_getcustomerbank_proto_rawDescData
}

var file_banking_aa_getcustomerbank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_banking_aa_getcustomerbank_proto_goTypes = []interface{}{
	(*GetCustomerBankRequest)(nil),  // 0: banking.aa.getcustomerbank.getCustomerBankRequest
	(*GetCustomerBankResponse)(nil), // 1: banking.aa.getcustomerbank.getCustomerBankResponse
	(*BankDetails)(nil),             // 2: banking.aa.getcustomerbank.BankDetails
}
var file_banking_aa_getcustomerbank_proto_depIdxs = []int32{
	2, // 0: banking.aa.getcustomerbank.getCustomerBankResponse.data:type_name -> banking.aa.getcustomerbank.BankDetails
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_banking_aa_getcustomerbank_proto_init() }
func file_banking_aa_getcustomerbank_proto_init() {
	if File_banking_aa_getcustomerbank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banking_aa_getcustomerbank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerBankRequest); i {
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
		file_banking_aa_getcustomerbank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerBankResponse); i {
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
		file_banking_aa_getcustomerbank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankDetails); i {
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
	file_banking_aa_getcustomerbank_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_banking_aa_getcustomerbank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banking_aa_getcustomerbank_proto_goTypes,
		DependencyIndexes: file_banking_aa_getcustomerbank_proto_depIdxs,
		MessageInfos:      file_banking_aa_getcustomerbank_proto_msgTypes,
	}.Build()
	File_banking_aa_getcustomerbank_proto = out.File
	file_banking_aa_getcustomerbank_proto_rawDesc = nil
	file_banking_aa_getcustomerbank_proto_goTypes = nil
	file_banking_aa_getcustomerbank_proto_depIdxs = nil
}
