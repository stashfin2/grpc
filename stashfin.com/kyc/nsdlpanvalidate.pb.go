// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: kyc/nsdlpanvalidate.proto

package kyc

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

type CustomerPanDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	FullName   string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	PanNumber  string `protobuf:"bytes,3,opt,name=pan_number,json=panNumber,proto3" json:"pan_number,omitempty"`
	Dob        string `protobuf:"bytes,4,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *CustomerPanDetailRequest) Reset() {
	*x = CustomerPanDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerPanDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerPanDetailRequest) ProtoMessage() {}

func (x *CustomerPanDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerPanDetailRequest.ProtoReflect.Descriptor instead.
func (*CustomerPanDetailRequest) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerPanDetailRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerPanDetailRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CustomerPanDetailRequest) GetPanNumber() string {
	if x != nil {
		return x.PanNumber
	}
	return ""
}

func (x *CustomerPanDetailRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

type InputData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pan        string  `protobuf:"bytes,1,opt,name=pan,proto3" json:"pan,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Fathername *string `protobuf:"bytes,3,opt,name=fathername,proto3,oneof" json:"fathername,omitempty"`
	Dob        string  `protobuf:"bytes,4,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *InputData) Reset() {
	*x = InputData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputData) ProtoMessage() {}

func (x *InputData) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputData.ProtoReflect.Descriptor instead.
func (*InputData) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{1}
}

func (x *InputData) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *InputData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InputData) GetFathername() string {
	if x != nil && x.Fathername != nil {
		return *x.Fathername
	}
	return ""
}

func (x *InputData) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

type NSDLRequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputData []*InputData `protobuf:"bytes,1,rep,name=inputData,proto3" json:"inputData,omitempty"`
	Signature string       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *NSDLRequestBody) Reset() {
	*x = NSDLRequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSDLRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSDLRequestBody) ProtoMessage() {}

func (x *NSDLRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSDLRequestBody.ProtoReflect.Descriptor instead.
func (*NSDLRequestBody) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{2}
}

func (x *NSDLRequestBody) GetInputData() []*InputData {
	if x != nil {
		return x.InputData
	}
	return nil
}

func (x *NSDLRequestBody) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type OutputData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pan           string  `protobuf:"bytes,1,opt,name=pan,proto3" json:"pan,omitempty"`
	PanStatus     *string `protobuf:"bytes,2,opt,name=pan_status,json=panStatus,proto3,oneof" json:"pan_status,omitempty"`
	Name          *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Fathername    *string `protobuf:"bytes,4,opt,name=fathername,proto3,oneof" json:"fathername,omitempty"`
	Dob           *string `protobuf:"bytes,5,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	SeedingStatus *string `protobuf:"bytes,6,opt,name=seeding_status,json=seedingStatus,proto3,oneof" json:"seeding_status,omitempty"`
}

func (x *OutputData) Reset() {
	*x = OutputData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputData) ProtoMessage() {}

func (x *OutputData) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputData.ProtoReflect.Descriptor instead.
func (*OutputData) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{3}
}

func (x *OutputData) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *OutputData) GetPanStatus() string {
	if x != nil && x.PanStatus != nil {
		return *x.PanStatus
	}
	return ""
}

func (x *OutputData) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *OutputData) GetFathername() string {
	if x != nil && x.Fathername != nil {
		return *x.Fathername
	}
	return ""
}

func (x *OutputData) GetDob() string {
	if x != nil && x.Dob != nil {
		return *x.Dob
	}
	return ""
}

func (x *OutputData) GetSeedingStatus() string {
	if x != nil && x.SeedingStatus != nil {
		return *x.SeedingStatus
	}
	return ""
}

type NSDLAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response_Code string        `protobuf:"bytes,1,opt,name=response_Code,json=responseCode,proto3" json:"response_Code,omitempty"`
	OutputData    []*OutputData `protobuf:"bytes,2,rep,name=outputData,proto3" json:"outputData,omitempty"`
}

func (x *NSDLAPIResponse) Reset() {
	*x = NSDLAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSDLAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSDLAPIResponse) ProtoMessage() {}

func (x *NSDLAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSDLAPIResponse.ProtoReflect.Descriptor instead.
func (*NSDLAPIResponse) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{4}
}

func (x *NSDLAPIResponse) GetResponse_Code() string {
	if x != nil {
		return x.Response_Code
	}
	return ""
}

func (x *NSDLAPIResponse) GetOutputData() []*OutputData {
	if x != nil {
		return x.OutputData
	}
	return nil
}

type CustomerPanDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *NSDLRequestBody `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *NSDLAPIResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CustomerPanDetailResponse) Reset() {
	*x = CustomerPanDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerPanDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerPanDetailResponse) ProtoMessage() {}

func (x *CustomerPanDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_nsdlpanvalidate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerPanDetailResponse.ProtoReflect.Descriptor instead.
func (*CustomerPanDetailResponse) Descriptor() ([]byte, []int) {
	return file_kyc_nsdlpanvalidate_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerPanDetailResponse) GetRequest() *NSDLRequestBody {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CustomerPanDetailResponse) GetResponse() *NSDLAPIResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_kyc_nsdlpanvalidate_proto protoreflect.FileDescriptor

var file_kyc_nsdlpanvalidate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x79, 0x63, 0x2f, 0x6e, 0x73, 0x64, 0x6c, 0x70, 0x61, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6b, 0x79, 0x63,
	0x2e, 0x6e, 0x73, 0x64, 0x6c, 0x70, 0x61, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x89, 0x01, 0x0a, 0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f,
	0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x22, 0x77, 0x0a, 0x09,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x0f, 0x4e, 0x53, 0x44, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x3c, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x79,
	0x63, 0x2e, 0x6e, 0x73, 0x64, 0x6c, 0x70, 0x61, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x85, 0x02, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x61, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x0e, 0x73, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x73, 0x65, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70,
	0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x6f, 0x62, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x65,
	0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x77, 0x0a, 0x0f,
	0x4e, 0x53, 0x44, 0x4c, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6e,
	0x73, 0x64, 0x6c, 0x70, 0x61, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x50, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6e, 0x73, 0x64, 0x6c, 0x70,
	0x61, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x53, 0x44, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6e, 0x73, 0x64, 0x6c,
	0x70, 0x61, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x53, 0x44, 0x4c,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kyc_nsdlpanvalidate_proto_rawDescOnce sync.Once
	file_kyc_nsdlpanvalidate_proto_rawDescData = file_kyc_nsdlpanvalidate_proto_rawDesc
)

func file_kyc_nsdlpanvalidate_proto_rawDescGZIP() []byte {
	file_kyc_nsdlpanvalidate_proto_rawDescOnce.Do(func() {
		file_kyc_nsdlpanvalidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyc_nsdlpanvalidate_proto_rawDescData)
	})
	return file_kyc_nsdlpanvalidate_proto_rawDescData
}

var file_kyc_nsdlpanvalidate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kyc_nsdlpanvalidate_proto_goTypes = []interface{}{
	(*CustomerPanDetailRequest)(nil),  // 0: kyc.nsdlpanvalidate.CustomerPanDetailRequest
	(*InputData)(nil),                 // 1: kyc.nsdlpanvalidate.InputData
	(*NSDLRequestBody)(nil),           // 2: kyc.nsdlpanvalidate.NSDLRequestBody
	(*OutputData)(nil),                // 3: kyc.nsdlpanvalidate.OutputData
	(*NSDLAPIResponse)(nil),           // 4: kyc.nsdlpanvalidate.NSDLAPIResponse
	(*CustomerPanDetailResponse)(nil), // 5: kyc.nsdlpanvalidate.CustomerPanDetailResponse
}
var file_kyc_nsdlpanvalidate_proto_depIdxs = []int32{
	1, // 0: kyc.nsdlpanvalidate.NSDLRequestBody.inputData:type_name -> kyc.nsdlpanvalidate.InputData
	3, // 1: kyc.nsdlpanvalidate.NSDLAPIResponse.outputData:type_name -> kyc.nsdlpanvalidate.OutputData
	2, // 2: kyc.nsdlpanvalidate.CustomerPanDetailResponse.request:type_name -> kyc.nsdlpanvalidate.NSDLRequestBody
	4, // 3: kyc.nsdlpanvalidate.CustomerPanDetailResponse.response:type_name -> kyc.nsdlpanvalidate.NSDLAPIResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kyc_nsdlpanvalidate_proto_init() }
func file_kyc_nsdlpanvalidate_proto_init() {
	if File_kyc_nsdlpanvalidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyc_nsdlpanvalidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerPanDetailRequest); i {
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
		file_kyc_nsdlpanvalidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputData); i {
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
		file_kyc_nsdlpanvalidate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSDLRequestBody); i {
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
		file_kyc_nsdlpanvalidate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputData); i {
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
		file_kyc_nsdlpanvalidate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSDLAPIResponse); i {
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
		file_kyc_nsdlpanvalidate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerPanDetailResponse); i {
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
	file_kyc_nsdlpanvalidate_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_kyc_nsdlpanvalidate_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kyc_nsdlpanvalidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kyc_nsdlpanvalidate_proto_goTypes,
		DependencyIndexes: file_kyc_nsdlpanvalidate_proto_depIdxs,
		MessageInfos:      file_kyc_nsdlpanvalidate_proto_msgTypes,
	}.Build()
	File_kyc_nsdlpanvalidate_proto = out.File
	file_kyc_nsdlpanvalidate_proto_rawDesc = nil
	file_kyc_nsdlpanvalidate_proto_goTypes = nil
	file_kyc_nsdlpanvalidate_proto_depIdxs = nil
}
