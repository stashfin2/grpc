// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: loans/sendemail.proto

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

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId int32 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_sendemail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_sendemail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_loans_sendemail_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgreementEmail string `protobuf:"bytes,1,opt,name=agreement_email,json=agreementEmail,proto3" json:"agreement_email,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_sendemail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_sendemail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_loans_sendemail_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailResponse) GetAgreementEmail() string {
	if x != nil {
		return x.AgreementEmail
	}
	return ""
}

var File_loans_sendemail_proto protoreflect.FileDescriptor

var file_loans_sendemail_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2b, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x6f, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_sendemail_proto_rawDescOnce sync.Once
	file_loans_sendemail_proto_rawDescData = file_loans_sendemail_proto_rawDesc
)

func file_loans_sendemail_proto_rawDescGZIP() []byte {
	file_loans_sendemail_proto_rawDescOnce.Do(func() {
		file_loans_sendemail_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_sendemail_proto_rawDescData)
	})
	return file_loans_sendemail_proto_rawDescData
}

var file_loans_sendemail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loans_sendemail_proto_goTypes = []interface{}{
	(*SendEmailRequest)(nil),  // 0: loans.sendemail.sendEmailRequest
	(*SendEmailResponse)(nil), // 1: loans.sendemail.sendEmailResponse
}
var file_loans_sendemail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loans_sendemail_proto_init() }
func file_loans_sendemail_proto_init() {
	if File_loans_sendemail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_sendemail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_loans_sendemail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
			RawDescriptor: file_loans_sendemail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_sendemail_proto_goTypes,
		DependencyIndexes: file_loans_sendemail_proto_depIdxs,
		MessageInfos:      file_loans_sendemail_proto_msgTypes,
	}.Build()
	File_loans_sendemail_proto = out.File
	file_loans_sendemail_proto_rawDesc = nil
	file_loans_sendemail_proto_goTypes = nil
	file_loans_sendemail_proto_depIdxs = nil
}
