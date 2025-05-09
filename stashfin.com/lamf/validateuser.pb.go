// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: lamf/validateuser.proto

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

type ValidateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  int32   `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ReqId       *int32  `protobuf:"varint,2,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	ClientRefNo *string `protobuf:"bytes,3,opt,name=clientRefNo,proto3,oneof" json:"clientRefNo,omitempty"`
	Pan         *string `protobuf:"bytes,4,opt,name=pan,proto3,oneof" json:"pan,omitempty"`
	Email       *string `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Mobile      *string `protobuf:"bytes,6,opt,name=mobile,proto3,oneof" json:"mobile,omitempty"`
	RetryCount  *int32  `protobuf:"varint,7,opt,name=retryCount,proto3,oneof" json:"retryCount,omitempty"`
}

func (x *ValidateUserRequest) Reset() {
	*x = ValidateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_validateuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserRequest) ProtoMessage() {}

func (x *ValidateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserRequest.ProtoReflect.Descriptor instead.
func (*ValidateUserRequest) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateUserRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ValidateUserRequest) GetReqId() int32 {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return 0
}

func (x *ValidateUserRequest) GetClientRefNo() string {
	if x != nil && x.ClientRefNo != nil {
		return *x.ClientRefNo
	}
	return ""
}

func (x *ValidateUserRequest) GetPan() string {
	if x != nil && x.Pan != nil {
		return *x.Pan
	}
	return ""
}

func (x *ValidateUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *ValidateUserRequest) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *ValidateUserRequest) GetRetryCount() int32 {
	if x != nil && x.RetryCount != nil {
		return *x.RetryCount
	}
	return 0
}

type ValidateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       *ValidateUserResponse_Data `protobuf:"bytes,1,opt,name=data,proto3,oneof" json:"data,omitempty"`
	StatusCode *int32                     `protobuf:"varint,2,opt,name=statusCode,proto3,oneof" json:"statusCode,omitempty"`
}

func (x *ValidateUserResponse) Reset() {
	*x = ValidateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_validateuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserResponse) ProtoMessage() {}

func (x *ValidateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserResponse.ProtoReflect.Descriptor instead.
func (*ValidateUserResponse) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateUserResponse) GetData() *ValidateUserResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ValidateUserResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

type ValidateUserResponse_Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *int32  `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *ValidateUserResponse_Errors) Reset() {
	*x = ValidateUserResponse_Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_validateuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserResponse_Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserResponse_Errors) ProtoMessage() {}

func (x *ValidateUserResponse_Errors) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserResponse_Errors.ProtoReflect.Descriptor instead.
func (*ValidateUserResponse_Errors) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ValidateUserResponse_Errors) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *ValidateUserResponse_Errors) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ValidateUserResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId                *string                        `protobuf:"bytes,1,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	OtpRef               *string                        `protobuf:"bytes,2,opt,name=otpRef,proto3,oneof" json:"otpRef,omitempty"`
	UserSubjectReference *string                        `protobuf:"bytes,3,opt,name=userSubjectReference,proto3,oneof" json:"userSubjectReference,omitempty"`
	ClientRefNo          *string                        `protobuf:"bytes,4,opt,name=clientRefNo,proto3,oneof" json:"clientRefNo,omitempty"`
	Errors               []*ValidateUserResponse_Errors `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	AllowRetry           *string                        `protobuf:"bytes,6,opt,name=allowRetry,proto3,oneof" json:"allowRetry,omitempty"`
}

func (x *ValidateUserResponse_Data) Reset() {
	*x = ValidateUserResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_validateuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUserResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUserResponse_Data) ProtoMessage() {}

func (x *ValidateUserResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_validateuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUserResponse_Data.ProtoReflect.Descriptor instead.
func (*ValidateUserResponse_Data) Descriptor() ([]byte, []int) {
	return file_lamf_validateuser_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ValidateUserResponse_Data) GetReqId() string {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return ""
}

func (x *ValidateUserResponse_Data) GetOtpRef() string {
	if x != nil && x.OtpRef != nil {
		return *x.OtpRef
	}
	return ""
}

func (x *ValidateUserResponse_Data) GetUserSubjectReference() string {
	if x != nil && x.UserSubjectReference != nil {
		return *x.UserSubjectReference
	}
	return ""
}

func (x *ValidateUserResponse_Data) GetClientRefNo() string {
	if x != nil && x.ClientRefNo != nil {
		return *x.ClientRefNo
	}
	return ""
}

func (x *ValidateUserResponse_Data) GetErrors() []*ValidateUserResponse_Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ValidateUserResponse_Data) GetAllowRetry() string {
	if x != nil && x.AllowRetry != nil {
		return *x.AllowRetry
	}
	return ""
}

var File_lamf_validateuser_proto protoreflect.FileDescriptor

var file_lamf_validateuser_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x61, 0x6d, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6c, 0x61, 0x6d, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73, 0x65, 0x72, 0x22, 0xb2, 0x02, 0x0a,
	0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x66, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x70, 0x61, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x66, 0x4e, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x61, 0x6e, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xcc, 0x04, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x55, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xd8, 0x02,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x74, 0x70, 0x52, 0x65, 0x66, 0x88, 0x01, 0x01, 0x12, 0x37,
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x14,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x46,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
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
}

var (
	file_lamf_validateuser_proto_rawDescOnce sync.Once
	file_lamf_validateuser_proto_rawDescData = file_lamf_validateuser_proto_rawDesc
)

func file_lamf_validateuser_proto_rawDescGZIP() []byte {
	file_lamf_validateuser_proto_rawDescOnce.Do(func() {
		file_lamf_validateuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_lamf_validateuser_proto_rawDescData)
	})
	return file_lamf_validateuser_proto_rawDescData
}

var file_lamf_validateuser_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lamf_validateuser_proto_goTypes = []interface{}{
	(*ValidateUserRequest)(nil),         // 0: lamf.validateuser.validateUserRequest
	(*ValidateUserResponse)(nil),        // 1: lamf.validateuser.validateUserResponse
	(*ValidateUserResponse_Errors)(nil), // 2: lamf.validateuser.validateUserResponse.Errors
	(*ValidateUserResponse_Data)(nil),   // 3: lamf.validateuser.validateUserResponse.Data
}
var file_lamf_validateuser_proto_depIdxs = []int32{
	3, // 0: lamf.validateuser.validateUserResponse.data:type_name -> lamf.validateuser.validateUserResponse.Data
	2, // 1: lamf.validateuser.validateUserResponse.Data.errors:type_name -> lamf.validateuser.validateUserResponse.Errors
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
	if !protoimpl.UnsafeEnabled {
		file_lamf_validateuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserRequest); i {
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
		file_lamf_validateuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserResponse); i {
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
		file_lamf_validateuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserResponse_Errors); i {
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
		file_lamf_validateuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUserResponse_Data); i {
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
	file_lamf_validateuser_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_lamf_validateuser_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_lamf_validateuser_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_lamf_validateuser_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lamf_validateuser_proto_rawDesc,
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
	file_lamf_validateuser_proto_rawDesc = nil
	file_lamf_validateuser_proto_goTypes = nil
	file_lamf_validateuser_proto_depIdxs = nil
}
