// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: loans/banklist.proto

package loans

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

type BankListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BankListRequest) Reset() {
	*x = BankListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_banklist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankListRequest) ProtoMessage() {}

func (x *BankListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_banklist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankListRequest.ProtoReflect.Descriptor instead.
func (*BankListRequest) Descriptor() ([]byte, []int) {
	return file_loans_banklist_proto_rawDescGZIP(), []int{0}
}

type BankListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountNumber string `protobuf:"bytes,2,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	BankName      string `protobuf:"bytes,3,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	IfscCode      string `protobuf:"bytes,4,opt,name=ifsc_code,json=ifscCode,proto3" json:"ifsc_code,omitempty"`
}

func (x *BankListResponse) Reset() {
	*x = BankListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_banklist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankListResponse) ProtoMessage() {}

func (x *BankListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_banklist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankListResponse.ProtoReflect.Descriptor instead.
func (*BankListResponse) Descriptor() ([]byte, []int) {
	return file_loans_banklist_proto_rawDescGZIP(), []int{1}
}

func (x *BankListResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BankListResponse) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *BankListResponse) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *BankListResponse) GetIfscCode() string {
	if x != nil {
		return x.IfscCode
	}
	return ""
}

var File_loans_banklist_proto protoreflect.FileDescriptor

var file_loans_banklist_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x62, 0x61,
	0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x66, 0x73, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x66, 0x73, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_banklist_proto_rawDescOnce sync.Once
	file_loans_banklist_proto_rawDescData = file_loans_banklist_proto_rawDesc
)

func file_loans_banklist_proto_rawDescGZIP() []byte {
	file_loans_banklist_proto_rawDescOnce.Do(func() {
		file_loans_banklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_banklist_proto_rawDescData)
	})
	return file_loans_banklist_proto_rawDescData
}

var file_loans_banklist_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loans_banklist_proto_goTypes = []interface{}{
	(*BankListRequest)(nil),  // 0: loans.banklist.bankListRequest
	(*BankListResponse)(nil), // 1: loans.banklist.bankListResponse
}
var file_loans_banklist_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loans_banklist_proto_init() }
func file_loans_banklist_proto_init() {
	if File_loans_banklist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_banklist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankListRequest); i {
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
		file_loans_banklist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankListResponse); i {
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
			RawDescriptor: file_loans_banklist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_banklist_proto_goTypes,
		DependencyIndexes: file_loans_banklist_proto_depIdxs,
		MessageInfos:      file_loans_banklist_proto_msgTypes,
	}.Build()
	File_loans_banklist_proto = out.File
	file_loans_banklist_proto_rawDesc = nil
	file_loans_banklist_proto_goTypes = nil
	file_loans_banklist_proto_depIdxs = nil
}
