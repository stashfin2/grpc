// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: upi/checkderegisterpermission.proto

package upi

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

type CheckDeregisterPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile    string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	DeviceId  string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ProfileId string `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Vpa       string `protobuf:"bytes,4,opt,name=vpa,proto3" json:"vpa,omitempty"`
}

func (x *CheckDeregisterPermissionRequest) Reset() {
	*x = CheckDeregisterPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_checkderegisterpermission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDeregisterPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDeregisterPermissionRequest) ProtoMessage() {}

func (x *CheckDeregisterPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_checkderegisterpermission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDeregisterPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckDeregisterPermissionRequest) Descriptor() ([]byte, []int) {
	return file_upi_checkderegisterpermission_proto_rawDescGZIP(), []int{0}
}

func (x *CheckDeregisterPermissionRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CheckDeregisterPermissionRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *CheckDeregisterPermissionRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *CheckDeregisterPermissionRequest) GetVpa() string {
	if x != nil {
		return x.Vpa
	}
	return ""
}

type CheckDeregisterPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *CheckDeregisterPermissionResponse) Reset() {
	*x = CheckDeregisterPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upi_checkderegisterpermission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDeregisterPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDeregisterPermissionResponse) ProtoMessage() {}

func (x *CheckDeregisterPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_checkderegisterpermission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDeregisterPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckDeregisterPermissionResponse) Descriptor() ([]byte, []int) {
	return file_upi_checkderegisterpermission_proto_rawDescGZIP(), []int{1}
}

func (x *CheckDeregisterPermissionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CheckDeregisterPermissionResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_upi_checkderegisterpermission_proto protoreflect.FileDescriptor

var file_upi_checkderegisterpermission_proto_rawDesc = []byte{
	0x0a, 0x23, 0x75, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x64, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x75, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x70, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x70, 0x61, 0x22,
	0x66, 0x0a, 0x21, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upi_checkderegisterpermission_proto_rawDescOnce sync.Once
	file_upi_checkderegisterpermission_proto_rawDescData = file_upi_checkderegisterpermission_proto_rawDesc
)

func file_upi_checkderegisterpermission_proto_rawDescGZIP() []byte {
	file_upi_checkderegisterpermission_proto_rawDescOnce.Do(func() {
		file_upi_checkderegisterpermission_proto_rawDescData = protoimpl.X.CompressGZIP(file_upi_checkderegisterpermission_proto_rawDescData)
	})
	return file_upi_checkderegisterpermission_proto_rawDescData
}

var file_upi_checkderegisterpermission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_upi_checkderegisterpermission_proto_goTypes = []interface{}{
	(*CheckDeregisterPermissionRequest)(nil),  // 0: upi.checkDeregisterPermission.checkDeregisterPermissionRequest
	(*CheckDeregisterPermissionResponse)(nil), // 1: upi.checkDeregisterPermission.checkDeregisterPermissionResponse
}
var file_upi_checkderegisterpermission_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upi_checkderegisterpermission_proto_init() }
func file_upi_checkderegisterpermission_proto_init() {
	if File_upi_checkderegisterpermission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upi_checkderegisterpermission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDeregisterPermissionRequest); i {
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
		file_upi_checkderegisterpermission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDeregisterPermissionResponse); i {
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
	file_upi_checkderegisterpermission_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upi_checkderegisterpermission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_checkderegisterpermission_proto_goTypes,
		DependencyIndexes: file_upi_checkderegisterpermission_proto_depIdxs,
		MessageInfos:      file_upi_checkderegisterpermission_proto_msgTypes,
	}.Build()
	File_upi_checkderegisterpermission_proto = out.File
	file_upi_checkderegisterpermission_proto_rawDesc = nil
	file_upi_checkderegisterpermission_proto_goTypes = nil
	file_upi_checkderegisterpermission_proto_depIdxs = nil
}
