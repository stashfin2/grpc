// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: lamf/submitotp.proto

package lamf

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

type SubmitotpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int32                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Otp           string                 `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
	ReqId         string                 `protobuf:"bytes,3,opt,name=reqId,proto3" json:"reqId,omitempty"`
	OtpRef        string                 `protobuf:"bytes,4,opt,name=otpRef,proto3" json:"otpRef,omitempty"`
	ClientRefNo   string                 `protobuf:"bytes,5,opt,name=clientRefNo,proto3" json:"clientRefNo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitotpRequest) Reset() {
	*x = SubmitotpRequest{}
	mi := &file_lamf_submitotp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitotpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitotpRequest) ProtoMessage() {}

func (x *SubmitotpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_submitotp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitotpRequest.ProtoReflect.Descriptor instead.
func (*SubmitotpRequest) Descriptor() ([]byte, []int) {
	return file_lamf_submitotp_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitotpRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SubmitotpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *SubmitotpRequest) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *SubmitotpRequest) GetOtpRef() string {
	if x != nil {
		return x.OtpRef
	}
	return ""
}

func (x *SubmitotpRequest) GetClientRefNo() string {
	if x != nil {
		return x.ClientRefNo
	}
	return ""
}

type SubmitotpResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	StatusCode    *int32                  `protobuf:"varint,1,opt,name=statusCode,proto3,oneof" json:"statusCode,omitempty"`
	Data          *SubmitotpResponse_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitotpResponse) Reset() {
	*x = SubmitotpResponse{}
	mi := &file_lamf_submitotp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitotpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitotpResponse) ProtoMessage() {}

func (x *SubmitotpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_submitotp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitotpResponse.ProtoReflect.Descriptor instead.
func (*SubmitotpResponse) Descriptor() ([]byte, []int) {
	return file_lamf_submitotp_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitotpResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *SubmitotpResponse) GetData() *SubmitotpResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubmitotpResponse_Errors struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          *int32                 `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitotpResponse_Errors) Reset() {
	*x = SubmitotpResponse_Errors{}
	mi := &file_lamf_submitotp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitotpResponse_Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitotpResponse_Errors) ProtoMessage() {}

func (x *SubmitotpResponse_Errors) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_submitotp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitotpResponse_Errors.ProtoReflect.Descriptor instead.
func (*SubmitotpResponse_Errors) Descriptor() ([]byte, []int) {
	return file_lamf_submitotp_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SubmitotpResponse_Errors) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *SubmitotpResponse_Errors) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type SubmitotpResponse_Data struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	ReqId         *string                     `protobuf:"bytes,1,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	Errors        []*SubmitotpResponse_Errors `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitotpResponse_Data) Reset() {
	*x = SubmitotpResponse_Data{}
	mi := &file_lamf_submitotp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitotpResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitotpResponse_Data) ProtoMessage() {}

func (x *SubmitotpResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_submitotp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitotpResponse_Data.ProtoReflect.Descriptor instead.
func (*SubmitotpResponse_Data) Descriptor() ([]byte, []int) {
	return file_lamf_submitotp_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SubmitotpResponse_Data) GetReqId() string {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return ""
}

func (x *SubmitotpResponse_Data) GetErrors() []*SubmitotpResponse_Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_lamf_submitotp_proto protoreflect.FileDescriptor

var file_lamf_submitotp_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x6c, 0x61, 0x6d, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x6f, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x6f, 0x74, 0x70, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x22, 0xd7,
	0x02, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x6f, 0x74, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x6f,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x1a, 0x55, 0x0a, 0x06, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x6d, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x6f, 0x74, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x6f, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x71, 0x49, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_lamf_submitotp_proto_rawDescOnce sync.Once
	file_lamf_submitotp_proto_rawDescData []byte
)

func file_lamf_submitotp_proto_rawDescGZIP() []byte {
	file_lamf_submitotp_proto_rawDescOnce.Do(func() {
		file_lamf_submitotp_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lamf_submitotp_proto_rawDesc), len(file_lamf_submitotp_proto_rawDesc)))
	})
	return file_lamf_submitotp_proto_rawDescData
}

var file_lamf_submitotp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lamf_submitotp_proto_goTypes = []any{
	(*SubmitotpRequest)(nil),         // 0: lamf.submitotp.submitotpRequest
	(*SubmitotpResponse)(nil),        // 1: lamf.submitotp.submitotpResponse
	(*SubmitotpResponse_Errors)(nil), // 2: lamf.submitotp.submitotpResponse.Errors
	(*SubmitotpResponse_Data)(nil),   // 3: lamf.submitotp.submitotpResponse.Data
}
var file_lamf_submitotp_proto_depIdxs = []int32{
	3, // 0: lamf.submitotp.submitotpResponse.data:type_name -> lamf.submitotp.submitotpResponse.Data
	2, // 1: lamf.submitotp.submitotpResponse.Data.errors:type_name -> lamf.submitotp.submitotpResponse.Errors
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lamf_submitotp_proto_init() }
func file_lamf_submitotp_proto_init() {
	if File_lamf_submitotp_proto != nil {
		return
	}
	file_lamf_submitotp_proto_msgTypes[1].OneofWrappers = []any{}
	file_lamf_submitotp_proto_msgTypes[2].OneofWrappers = []any{}
	file_lamf_submitotp_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lamf_submitotp_proto_rawDesc), len(file_lamf_submitotp_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lamf_submitotp_proto_goTypes,
		DependencyIndexes: file_lamf_submitotp_proto_depIdxs,
		MessageInfos:      file_lamf_submitotp_proto_msgTypes,
	}.Build()
	File_lamf_submitotp_proto = out.File
	file_lamf_submitotp_proto_goTypes = nil
	file_lamf_submitotp_proto_depIdxs = nil
}
