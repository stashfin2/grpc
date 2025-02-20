// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: customers/profilebasic.proto

package customers

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

type UpdateProfileBasicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language   string  `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	AppVersion string  `protobuf:"bytes,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	DeviceId   string  `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	OsVersion  string  `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Longitude  float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude   float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	PushId     *string `protobuf:"bytes,7,opt,name=push_id,json=pushId,proto3,oneof" json:"push_id,omitempty"`
}

func (x *UpdateProfileBasicRequest) Reset() {
	*x = UpdateProfileBasicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_profilebasic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileBasicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileBasicRequest) ProtoMessage() {}

func (x *UpdateProfileBasicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_profilebasic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileBasicRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileBasicRequest) Descriptor() ([]byte, []int) {
	return file_customers_profilebasic_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateProfileBasicRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UpdateProfileBasicRequest) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *UpdateProfileBasicRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UpdateProfileBasicRequest) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *UpdateProfileBasicRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UpdateProfileBasicRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateProfileBasicRequest) GetPushId() string {
	if x != nil && x.PushId != nil {
		return *x.PushId
	}
	return ""
}

type UpdateProfileBasicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateProfileBasicResponse) Reset() {
	*x = UpdateProfileBasicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_profilebasic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileBasicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileBasicResponse) ProtoMessage() {}

func (x *UpdateProfileBasicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_profilebasic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileBasicResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileBasicResponse) Descriptor() ([]byte, []int) {
	return file_customers_profilebasic_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProfileBasicResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_customers_profilebasic_proto protoreflect.FileDescriptor

var file_customers_profilebasic_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0xf8, 0x01, 0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x75, 0x73, 0x68,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x69,
	0x64, 0x22, 0x34, 0x0a, 0x1a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_profilebasic_proto_rawDescOnce sync.Once
	file_customers_profilebasic_proto_rawDescData = file_customers_profilebasic_proto_rawDesc
)

func file_customers_profilebasic_proto_rawDescGZIP() []byte {
	file_customers_profilebasic_proto_rawDescOnce.Do(func() {
		file_customers_profilebasic_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_profilebasic_proto_rawDescData)
	})
	return file_customers_profilebasic_proto_rawDescData
}

var file_customers_profilebasic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_profilebasic_proto_goTypes = []interface{}{
	(*UpdateProfileBasicRequest)(nil),  // 0: customers.profilebasic.updateProfileBasicRequest
	(*UpdateProfileBasicResponse)(nil), // 1: customers.profilebasic.updateProfileBasicResponse
}
var file_customers_profilebasic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_profilebasic_proto_init() }
func file_customers_profilebasic_proto_init() {
	if File_customers_profilebasic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_profilebasic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileBasicRequest); i {
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
		file_customers_profilebasic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileBasicResponse); i {
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
	file_customers_profilebasic_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customers_profilebasic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_profilebasic_proto_goTypes,
		DependencyIndexes: file_customers_profilebasic_proto_depIdxs,
		MessageInfos:      file_customers_profilebasic_proto_msgTypes,
	}.Build()
	File_customers_profilebasic_proto = out.File
	file_customers_profilebasic_proto_rawDesc = nil
	file_customers_profilebasic_proto_goTypes = nil
	file_customers_profilebasic_proto_depIdxs = nil
}
