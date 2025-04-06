// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/validatevpa.proto

package upi

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

type ValidateVPARequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	DeviceId      string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ProfileId     string                 `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Vpa           string                 `protobuf:"bytes,4,opt,name=vpa,proto3" json:"vpa,omitempty"`
	ValidateType  *string                `protobuf:"bytes,5,opt,name=validate_type,json=validateType,proto3,oneof" json:"validate_type,omitempty"`
	Ifsc          *string                `protobuf:"bytes,6,opt,name=ifsc,proto3,oneof" json:"ifsc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateVPARequest) Reset() {
	*x = ValidateVPARequest{}
	mi := &file_upi_validatevpa_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateVPARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateVPARequest) ProtoMessage() {}

func (x *ValidateVPARequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_validatevpa_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateVPARequest.ProtoReflect.Descriptor instead.
func (*ValidateVPARequest) Descriptor() ([]byte, []int) {
	return file_upi_validatevpa_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateVPARequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ValidateVPARequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ValidateVPARequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *ValidateVPARequest) GetVpa() string {
	if x != nil {
		return x.Vpa
	}
	return ""
}

func (x *ValidateVPARequest) GetValidateType() string {
	if x != nil && x.ValidateType != nil {
		return *x.ValidateType
	}
	return ""
}

func (x *ValidateVPARequest) GetIfsc() string {
	if x != nil && x.Ifsc != nil {
		return *x.Ifsc
	}
	return ""
}

type ValidateVPAResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vpa           *string                `protobuf:"bytes,1,opt,name=vpa,proto3,oneof" json:"vpa,omitempty"`
	VpaData       *VpaData               `protobuf:"bytes,2,opt,name=vpaData,proto3,oneof" json:"vpaData,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateVPAResponse) Reset() {
	*x = ValidateVPAResponse{}
	mi := &file_upi_validatevpa_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateVPAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateVPAResponse) ProtoMessage() {}

func (x *ValidateVPAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_validatevpa_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateVPAResponse.ProtoReflect.Descriptor instead.
func (*ValidateVPAResponse) Descriptor() ([]byte, []int) {
	return file_upi_validatevpa_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateVPAResponse) GetVpa() string {
	if x != nil && x.Vpa != nil {
		return *x.Vpa
	}
	return ""
}

func (x *ValidateVPAResponse) GetVpaData() *VpaData {
	if x != nil {
		return x.VpaData
	}
	return nil
}

type VpaData struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	MaskName         *string                `protobuf:"bytes,1,opt,name=maskName,proto3,oneof" json:"maskName,omitempty"`
	Type             *string                `protobuf:"bytes,2,opt,name=type,proto3,oneof" json:"type,omitempty"`
	Code             *string                `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Mid              *string                `protobuf:"bytes,4,opt,name=mid,proto3,oneof" json:"mid,omitempty"`
	Sid              *string                `protobuf:"bytes,5,opt,name=sid,proto3,oneof" json:"sid,omitempty"`
	Tid              *string                `protobuf:"bytes,6,opt,name=tid,proto3,oneof" json:"tid,omitempty"`
	MerchantGenre    *string                `protobuf:"bytes,7,opt,name=merchantGenre,proto3,oneof" json:"merchantGenre,omitempty"`
	OnBoardingType   *string                `protobuf:"bytes,8,opt,name=onBoardingType,proto3,oneof" json:"onBoardingType,omitempty"`
	VerifiedMerchant *bool                  `protobuf:"varint,9,opt,name=verifiedMerchant,proto3,oneof" json:"verifiedMerchant,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *VpaData) Reset() {
	*x = VpaData{}
	mi := &file_upi_validatevpa_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VpaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VpaData) ProtoMessage() {}

func (x *VpaData) ProtoReflect() protoreflect.Message {
	mi := &file_upi_validatevpa_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VpaData.ProtoReflect.Descriptor instead.
func (*VpaData) Descriptor() ([]byte, []int) {
	return file_upi_validatevpa_proto_rawDescGZIP(), []int{2}
}

func (x *VpaData) GetMaskName() string {
	if x != nil && x.MaskName != nil {
		return *x.MaskName
	}
	return ""
}

func (x *VpaData) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *VpaData) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *VpaData) GetMid() string {
	if x != nil && x.Mid != nil {
		return *x.Mid
	}
	return ""
}

func (x *VpaData) GetSid() string {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return ""
}

func (x *VpaData) GetTid() string {
	if x != nil && x.Tid != nil {
		return *x.Tid
	}
	return ""
}

func (x *VpaData) GetMerchantGenre() string {
	if x != nil && x.MerchantGenre != nil {
		return *x.MerchantGenre
	}
	return ""
}

func (x *VpaData) GetOnBoardingType() string {
	if x != nil && x.OnBoardingType != nil {
		return *x.OnBoardingType
	}
	return ""
}

func (x *VpaData) GetVerifiedMerchant() bool {
	if x != nil && x.VerifiedMerchant != nil {
		return *x.VerifiedMerchant
	}
	return false
}

var File_upi_validatevpa_proto protoreflect.FileDescriptor

var file_upi_validatevpa_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x75, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x76, 0x70,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x50, 0x41, 0x22, 0xd8, 0x01, 0x0a, 0x12, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x50, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x70, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x76, 0x70, 0x61, 0x12, 0x28, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x69, 0x66, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x69, 0x66, 0x73, 0x63, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69,
	0x66, 0x73, 0x63, 0x22, 0x79, 0x0a, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x50, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x76, 0x70,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x76, 0x70, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x37, 0x0a, 0x07, 0x76, 0x70, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x50, 0x41, 0x2e, 0x56, 0x70, 0x61, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x07,
	0x76, 0x70, 0x61, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76,
	0x70, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x70, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9b,
	0x03, 0x0a, 0x07, 0x56, 0x70, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x61,
	0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x6d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x03, 0x6d, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x03, 0x73, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x74,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0d, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a,
	0x0e, 0x6f, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0e, 0x6f, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x08, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d,
	0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74,
	0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_validatevpa_proto_rawDescOnce sync.Once
	file_upi_validatevpa_proto_rawDescData []byte
)

func file_upi_validatevpa_proto_rawDescGZIP() []byte {
	file_upi_validatevpa_proto_rawDescOnce.Do(func() {
		file_upi_validatevpa_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_validatevpa_proto_rawDesc), len(file_upi_validatevpa_proto_rawDesc)))
	})
	return file_upi_validatevpa_proto_rawDescData
}

var file_upi_validatevpa_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_upi_validatevpa_proto_goTypes = []any{
	(*ValidateVPARequest)(nil),  // 0: upi.validateVPA.validateVPARequest
	(*ValidateVPAResponse)(nil), // 1: upi.validateVPA.validateVPAResponse
	(*VpaData)(nil),             // 2: upi.validateVPA.VpaData
}
var file_upi_validatevpa_proto_depIdxs = []int32{
	2, // 0: upi.validateVPA.validateVPAResponse.vpaData:type_name -> upi.validateVPA.VpaData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_upi_validatevpa_proto_init() }
func file_upi_validatevpa_proto_init() {
	if File_upi_validatevpa_proto != nil {
		return
	}
	file_upi_validatevpa_proto_msgTypes[0].OneofWrappers = []any{}
	file_upi_validatevpa_proto_msgTypes[1].OneofWrappers = []any{}
	file_upi_validatevpa_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_validatevpa_proto_rawDesc), len(file_upi_validatevpa_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_validatevpa_proto_goTypes,
		DependencyIndexes: file_upi_validatevpa_proto_depIdxs,
		MessageInfos:      file_upi_validatevpa_proto_msgTypes,
	}.Build()
	File_upi_validatevpa_proto = out.File
	file_upi_validatevpa_proto_goTypes = nil
	file_upi_validatevpa_proto_depIdxs = nil
}
