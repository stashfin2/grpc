// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: lamf/validateuser.proto

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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int32                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ReqId         *int32                 `protobuf:"varint,2,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	ClientRefNo   *string                `protobuf:"bytes,3,opt,name=clientRefNo,proto3,oneof" json:"clientRefNo,omitempty"`
	Pan           *string                `protobuf:"bytes,4,opt,name=pan,proto3,oneof" json:"pan,omitempty"`
	Email         *string                `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Mobile        *string                `protobuf:"bytes,6,opt,name=mobile,proto3,oneof" json:"mobile,omitempty"`
	RetryCount    *int32                 `protobuf:"varint,7,opt,name=retryCount,proto3,oneof" json:"retryCount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_lamf_validateuser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Request) GetReqId() int32 {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return 0
}

func (x *Request) GetClientRefNo() string {
	if x != nil && x.ClientRefNo != nil {
		return *x.ClientRefNo
	}
	return ""
}

func (x *Request) GetPan() string {
	if x != nil && x.Pan != nil {
		return *x.Pan
	}
	return ""
}

func (x *Request) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Request) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *Request) GetRetryCount() int32 {
	if x != nil && x.RetryCount != nil {
		return *x.RetryCount
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Response_Data         `protobuf:"bytes,1,opt,name=data,proto3,oneof" json:"data,omitempty"`
	StatusCode    *int32                 `protobuf:"varint,2,opt,name=statusCode,proto3,oneof" json:"statusCode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_lamf_validateuser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() *Response_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

type Response_Errors struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          *int32                 `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response_Errors) Reset() {
	*x = Response_Errors{}
	mi := &file_lamf_validateuser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response_Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Errors) ProtoMessage() {}

func (x *Response_Errors) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Errors.ProtoReflect.Descriptor instead.
func (*Response_Errors) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_Errors) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *Response_Errors) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type Response_Data struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	ReqId                *string                `protobuf:"bytes,1,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	OtpRef               *string                `protobuf:"bytes,2,opt,name=otpRef,proto3,oneof" json:"otpRef,omitempty"`
	UserSubjectReference *string                `protobuf:"bytes,3,opt,name=userSubjectReference,proto3,oneof" json:"userSubjectReference,omitempty"`
	ClientRefNo          *string                `protobuf:"bytes,4,opt,name=clientRefNo,proto3,oneof" json:"clientRefNo,omitempty"`
	Errors               []*Response_Errors     `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	AllowRetry           *string                `protobuf:"bytes,6,opt,name=allowRetry,proto3,oneof" json:"allowRetry,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Response_Data) Reset() {
	*x = Response_Data{}
	mi := &file_lamf_validateuser_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Data) ProtoMessage() {}

func (x *Response_Data) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Data.ProtoReflect.Descriptor instead.
func (*Response_Data) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Response_Data) GetReqId() string {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return ""
}

func (x *Response_Data) GetOtpRef() string {
	if x != nil && x.OtpRef != nil {
		return *x.OtpRef
	}
	return ""
}

func (x *Response_Data) GetUserSubjectReference() string {
	if x != nil && x.UserSubjectReference != nil {
		return *x.UserSubjectReference
	}
	return ""
}

func (x *Response_Data) GetClientRefNo() string {
	if x != nil && x.ClientRefNo != nil {
		return *x.ClientRefNo
	}
	return ""
}

func (x *Response_Data) GetErrors() []*Response_Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Response_Data) GetAllowRetry() string {
	if x != nil && x.AllowRetry != nil {
		return *x.AllowRetry
	}
	return ""
}

var File_lamf_validateuser_proto protoreflect.FileDescriptor

var file_lamf_validateuser_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x6c, 0x61, 0x6d, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6c, 0x61, 0x6d, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa6, 0x02, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x66, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x70,
	0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x70, 0x61, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05,
	0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x71, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x61,
	0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa8, 0x04, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x1a, 0x55, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xcc, 0x02, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x14, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x61, 0x6d,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66,
	0x42, 0x17, 0x0a, 0x15, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_lamf_validateuser_proto_rawDescOnce sync.Once
	file_lamf_validateuser_proto_rawDescData []byte
)

func file_lamf_validateuser_proto_rawDescGZIP() []byte {
	file_lamf_validateuser_proto_rawDescOnce.Do(func() {
		file_lamf_validateuser_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lamf_validateuser_proto_rawDesc), len(file_lamf_validateuser_proto_rawDesc)))
	})
	return file_lamf_validateuser_proto_rawDescData
}

var file_lamf_validateuser_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lamf_validateuser_proto_goTypes = []any{
	(*Request)(nil),         // 0: lamf.validateuser.request
	(*Response)(nil),        // 1: lamf.validateuser.response
	(*Response_Errors)(nil), // 2: lamf.validateuser.response.Errors
	(*Response_Data)(nil),   // 3: lamf.validateuser.response.Data
}
var file_lamf_validateuser_proto_depIdxs = []int32{
	3, // 0: lamf.validateuser.response.data:type_name -> lamf.validateuser.response.Data
	2, // 1: lamf.validateuser.response.Data.errors:type_name -> lamf.validateuser.response.Errors
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lamf_validateuser_proto_init() }
func file_lamf_validateuser_proto_init() {
	if File_lamf_validateuser_proto != nil {
		return
	}
	file_lamf_validateuser_proto_msgTypes[0].OneofWrappers = []any{}
	file_lamf_validateuser_proto_msgTypes[1].OneofWrappers = []any{}
	file_lamf_validateuser_proto_msgTypes[2].OneofWrappers = []any{}
	file_lamf_validateuser_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lamf_validateuser_proto_rawDesc), len(file_lamf_validateuser_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lamf_validateuser_proto_goTypes,
		DependencyIndexes: file_lamf_validateuser_proto_depIdxs,
		MessageInfos:      file_lamf_validateuser_proto_msgTypes,
	}.Build()
	File_lamf_validateuser_proto = out.File
	file_lamf_validateuser_proto_goTypes = nil
	file_lamf_validateuser_proto_depIdxs = nil
}
