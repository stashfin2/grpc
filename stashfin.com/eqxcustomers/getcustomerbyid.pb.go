// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: eqxcustomers/getcustomerbyid.proto

package eqxcustomers

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

type GetCustomerByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerByIdRequest) Reset() {
	*x = GetCustomerByIdRequest{}
	mi := &file_eqxcustomers_getcustomerbyid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdRequest) ProtoMessage() {}

func (x *GetCustomerByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getcustomerbyid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdRequest) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getcustomerbyid_proto_rawDescGZIP(), []int{0}
}

type GetCustomerByIdResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile         string                 `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status         string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	BiometricState int32                  `protobuf:"varint,4,opt,name=biometric_state,json=biometricState,proto3" json:"biometric_state,omitempty"`
	Vpa            *string                `protobuf:"bytes,5,opt,name=vpa,proto3,oneof" json:"vpa,omitempty"`
	FirstName      *string                `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3,oneof" json:"first_name,omitempty"`
	LastName       *string                `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3,oneof" json:"last_name,omitempty"`
	Email          *string                `protobuf:"bytes,8,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Token          *string                `protobuf:"bytes,9,opt,name=token,proto3,oneof" json:"token,omitempty"`
	RefferalCode   *string                `protobuf:"bytes,10,opt,name=refferal_code,json=refferalCode,proto3,oneof" json:"refferal_code,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetCustomerByIdResponse) Reset() {
	*x = GetCustomerByIdResponse{}
	mi := &file_eqxcustomers_getcustomerbyid_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdResponse) ProtoMessage() {}

func (x *GetCustomerByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_getcustomerbyid_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdResponse) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_getcustomerbyid_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerByIdResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCustomerByIdResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetBiometricState() int32 {
	if x != nil {
		return x.BiometricState
	}
	return 0
}

func (x *GetCustomerByIdResponse) GetVpa() string {
	if x != nil && x.Vpa != nil {
		return *x.Vpa
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *GetCustomerByIdResponse) GetRefferalCode() string {
	if x != nil && x.RefferalCode != nil {
		return *x.RefferalCode
	}
	return ""
}

var File_eqxcustomers_getcustomerbyid_proto protoreflect.FileDescriptor

var file_eqxcustomers_getcustomerbyid_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x79, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x79,
	0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8a, 0x03, 0x0a,
	0x17, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x69, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x62, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x03, 0x76, 0x70, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x76, 0x70, 0x61, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x66, 0x65, 0x72, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x66, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x76, 0x70, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x66, 0x66,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_eqxcustomers_getcustomerbyid_proto_rawDescOnce sync.Once
	file_eqxcustomers_getcustomerbyid_proto_rawDescData []byte
)

func file_eqxcustomers_getcustomerbyid_proto_rawDescGZIP() []byte {
	file_eqxcustomers_getcustomerbyid_proto_rawDescOnce.Do(func() {
		file_eqxcustomers_getcustomerbyid_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eqxcustomers_getcustomerbyid_proto_rawDesc), len(file_eqxcustomers_getcustomerbyid_proto_rawDesc)))
	})
	return file_eqxcustomers_getcustomerbyid_proto_rawDescData
}

var file_eqxcustomers_getcustomerbyid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eqxcustomers_getcustomerbyid_proto_goTypes = []any{
	(*GetCustomerByIdRequest)(nil),  // 0: eqxcustomers.getcustomerbyid.getCustomerByIdRequest
	(*GetCustomerByIdResponse)(nil), // 1: eqxcustomers.getcustomerbyid.getCustomerByIdResponse
}
var file_eqxcustomers_getcustomerbyid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eqxcustomers_getcustomerbyid_proto_init() }
func file_eqxcustomers_getcustomerbyid_proto_init() {
	if File_eqxcustomers_getcustomerbyid_proto != nil {
		return
	}
	file_eqxcustomers_getcustomerbyid_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eqxcustomers_getcustomerbyid_proto_rawDesc), len(file_eqxcustomers_getcustomerbyid_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eqxcustomers_getcustomerbyid_proto_goTypes,
		DependencyIndexes: file_eqxcustomers_getcustomerbyid_proto_depIdxs,
		MessageInfos:      file_eqxcustomers_getcustomerbyid_proto_msgTypes,
	}.Build()
	File_eqxcustomers_getcustomerbyid_proto = out.File
	file_eqxcustomers_getcustomerbyid_proto_goTypes = nil
	file_eqxcustomers_getcustomerbyid_proto_depIdxs = nil
}
