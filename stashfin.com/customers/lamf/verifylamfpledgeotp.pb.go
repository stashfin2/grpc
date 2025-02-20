// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: customers/lamf/verifylamfpledgeotp.proto

package lamf

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

type VerifyLamfPledgeOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *VerifyLamfPledgeOtpRequest) Reset() {
	*x = VerifyLamfPledgeOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyLamfPledgeOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyLamfPledgeOtpRequest) ProtoMessage() {}

func (x *VerifyLamfPledgeOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyLamfPledgeOtpRequest.ProtoReflect.Descriptor instead.
func (*VerifyLamfPledgeOtpRequest) Descriptor() ([]byte, []int) {
	return file_customers_lamf_verifylamfpledgeotp_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyLamfPledgeOtpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type VerifyLamfPledgeOtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VerifyLamfPledgeOtpResponse) Reset() {
	*x = VerifyLamfPledgeOtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyLamfPledgeOtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyLamfPledgeOtpResponse) ProtoMessage() {}

func (x *VerifyLamfPledgeOtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyLamfPledgeOtpResponse.ProtoReflect.Descriptor instead.
func (*VerifyLamfPledgeOtpResponse) Descriptor() ([]byte, []int) {
	return file_customers_lamf_verifylamfpledgeotp_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyLamfPledgeOtpResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_customers_lamf_verifylamfpledgeotp_proto protoreflect.FileDescriptor

var file_customers_lamf_verifylamfpledgeotp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x61, 0x6d, 0x66,
	0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x6c, 0x61, 0x6d, 0x66, 0x70, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x6c, 0x61, 0x6d, 0x66, 0x70, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6f, 0x74, 0x70, 0x22, 0x2e,
	0x0a, 0x1a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x61, 0x6d, 0x66, 0x50, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x35,
	0x0a, 0x1b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4c, 0x61, 0x6d, 0x66, 0x50, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_lamf_verifylamfpledgeotp_proto_rawDescOnce sync.Once
	file_customers_lamf_verifylamfpledgeotp_proto_rawDescData = file_customers_lamf_verifylamfpledgeotp_proto_rawDesc
)

func file_customers_lamf_verifylamfpledgeotp_proto_rawDescGZIP() []byte {
	file_customers_lamf_verifylamfpledgeotp_proto_rawDescOnce.Do(func() {
		file_customers_lamf_verifylamfpledgeotp_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_lamf_verifylamfpledgeotp_proto_rawDescData)
	})
	return file_customers_lamf_verifylamfpledgeotp_proto_rawDescData
}

var file_customers_lamf_verifylamfpledgeotp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_lamf_verifylamfpledgeotp_proto_goTypes = []interface{}{
	(*VerifyLamfPledgeOtpRequest)(nil),  // 0: customers.lamf.verifylamfpledgeotp.verifyLamfPledgeOtpRequest
	(*VerifyLamfPledgeOtpResponse)(nil), // 1: customers.lamf.verifylamfpledgeotp.verifyLamfPledgeOtpResponse
}
var file_customers_lamf_verifylamfpledgeotp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_lamf_verifylamfpledgeotp_proto_init() }
func file_customers_lamf_verifylamfpledgeotp_proto_init() {
	if File_customers_lamf_verifylamfpledgeotp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyLamfPledgeOtpRequest); i {
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
		file_customers_lamf_verifylamfpledgeotp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyLamfPledgeOtpResponse); i {
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
			RawDescriptor: file_customers_lamf_verifylamfpledgeotp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_lamf_verifylamfpledgeotp_proto_goTypes,
		DependencyIndexes: file_customers_lamf_verifylamfpledgeotp_proto_depIdxs,
		MessageInfos:      file_customers_lamf_verifylamfpledgeotp_proto_msgTypes,
	}.Build()
	File_customers_lamf_verifylamfpledgeotp_proto = out.File
	file_customers_lamf_verifylamfpledgeotp_proto_rawDesc = nil
	file_customers_lamf_verifylamfpledgeotp_proto_goTypes = nil
	file_customers_lamf_verifylamfpledgeotp_proto_depIdxs = nil
}
