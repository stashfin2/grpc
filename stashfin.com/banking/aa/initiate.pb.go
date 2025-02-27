// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: banking/aa/initiate.proto

package aa

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

type InitiateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JourneyType   string                 `protobuf:"bytes,1,opt,name=journey_type,json=journeyType,proto3" json:"journey_type,omitempty"`
	BankIds       []string               `protobuf:"bytes,2,rep,name=bank_ids,json=bankIds,proto3" json:"bank_ids,omitempty"`
	CustomerId    int64                  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitiateRequest) Reset() {
	*x = InitiateRequest{}
	mi := &file_banking_aa_initiate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitiateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateRequest) ProtoMessage() {}

func (x *InitiateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_initiate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateRequest.ProtoReflect.Descriptor instead.
func (*InitiateRequest) Descriptor() ([]byte, []int) {
	return file_banking_aa_initiate_proto_rawDescGZIP(), []int{0}
}

func (x *InitiateRequest) GetJourneyType() string {
	if x != nil {
		return x.JourneyType
	}
	return ""
}

func (x *InitiateRequest) GetBankIds() []string {
	if x != nil {
		return x.BankIds
	}
	return nil
}

func (x *InitiateRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type InitiateResponse struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Success                 *bool                  `protobuf:"varint,1,opt,name=success,proto3,oneof" json:"success,omitempty"`
	JourneyType             string                 `protobuf:"bytes,2,opt,name=journey_type,json=journeyType,proto3" json:"journey_type,omitempty"`
	MonitoringReferenceId   string                 `protobuf:"bytes,3,opt,name=monitoring_reference_id,json=monitoringReferenceId,proto3" json:"monitoring_reference_id,omitempty"`
	UnderwritingReferenceId string                 `protobuf:"bytes,4,opt,name=underwriting_reference_id,json=underwritingReferenceId,proto3" json:"underwriting_reference_id,omitempty"`
	RedirectionUrl          string                 `protobuf:"bytes,5,opt,name=redirection_url,json=redirectionUrl,proto3" json:"redirection_url,omitempty"`
	ReturnUrl               string                 `protobuf:"bytes,6,opt,name=return_url,json=returnUrl,proto3" json:"return_url,omitempty"`
	ErrorMessage            *string                `protobuf:"bytes,7,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
	ErrorCode               *string                `protobuf:"bytes,8,opt,name=error_code,json=errorCode,proto3,oneof" json:"error_code,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *InitiateResponse) Reset() {
	*x = InitiateResponse{}
	mi := &file_banking_aa_initiate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitiateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateResponse) ProtoMessage() {}

func (x *InitiateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_initiate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateResponse.ProtoReflect.Descriptor instead.
func (*InitiateResponse) Descriptor() ([]byte, []int) {
	return file_banking_aa_initiate_proto_rawDescGZIP(), []int{1}
}

func (x *InitiateResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *InitiateResponse) GetJourneyType() string {
	if x != nil {
		return x.JourneyType
	}
	return ""
}

func (x *InitiateResponse) GetMonitoringReferenceId() string {
	if x != nil {
		return x.MonitoringReferenceId
	}
	return ""
}

func (x *InitiateResponse) GetUnderwritingReferenceId() string {
	if x != nil {
		return x.UnderwritingReferenceId
	}
	return ""
}

func (x *InitiateResponse) GetRedirectionUrl() string {
	if x != nil {
		return x.RedirectionUrl
	}
	return ""
}

func (x *InitiateResponse) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

func (x *InitiateResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *InitiateResponse) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

var File_banking_aa_initiate_proto protoreflect.FileDescriptor

var file_banking_aa_initiate_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x22, 0x70, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x8b, 0x03, 0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_banking_aa_initiate_proto_rawDescOnce sync.Once
	file_banking_aa_initiate_proto_rawDescData []byte
)

func file_banking_aa_initiate_proto_rawDescGZIP() []byte {
	file_banking_aa_initiate_proto_rawDescOnce.Do(func() {
		file_banking_aa_initiate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_banking_aa_initiate_proto_rawDesc), len(file_banking_aa_initiate_proto_rawDesc)))
	})
	return file_banking_aa_initiate_proto_rawDescData
}

var file_banking_aa_initiate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_banking_aa_initiate_proto_goTypes = []any{
	(*InitiateRequest)(nil),  // 0: banking.aa.initiate.initiateRequest
	(*InitiateResponse)(nil), // 1: banking.aa.initiate.initiateResponse
}
var file_banking_aa_initiate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_banking_aa_initiate_proto_init() }
func file_banking_aa_initiate_proto_init() {
	if File_banking_aa_initiate_proto != nil {
		return
	}
	file_banking_aa_initiate_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_banking_aa_initiate_proto_rawDesc), len(file_banking_aa_initiate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banking_aa_initiate_proto_goTypes,
		DependencyIndexes: file_banking_aa_initiate_proto_depIdxs,
		MessageInfos:      file_banking_aa_initiate_proto_msgTypes,
	}.Build()
	File_banking_aa_initiate_proto = out.File
	file_banking_aa_initiate_proto_goTypes = nil
	file_banking_aa_initiate_proto_depIdxs = nil
}
