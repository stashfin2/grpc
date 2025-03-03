// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: eqxcustomers/sendotp.proto

package eqxcustomers

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

type SendOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *SendOtpRequest) Reset() {
	*x = SendOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_sendotp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpRequest) ProtoMessage() {}

func (x *SendOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_sendotp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpRequest.ProtoReflect.Descriptor instead.
func (*SendOtpRequest) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_sendotp_proto_rawDescGZIP(), []int{0}
}

func (x *SendOtpRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SendOtpRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type SendOtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IsRegistered bool   `protobuf:"varint,2,opt,name=is_registered,json=isRegistered,proto3" json:"is_registered,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendOtpResponse) Reset() {
	*x = SendOtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eqxcustomers_sendotp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpResponse) ProtoMessage() {}

func (x *SendOtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eqxcustomers_sendotp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpResponse.ProtoReflect.Descriptor instead.
func (*SendOtpResponse) Descriptor() ([]byte, []int) {
	return file_eqxcustomers_sendotp_proto_rawDescGZIP(), []int{1}
}

func (x *SendOtpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendOtpResponse) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

func (x *SendOtpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_eqxcustomers_sendotp_proto protoreflect.FileDescriptor

var file_eqxcustomers_sendotp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x71, 0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x71,
	0x78, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x6f,
	0x74, 0x70, 0x22, 0x45, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0f, 0x73, 0x65, 0x6e,
	0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eqxcustomers_sendotp_proto_rawDescOnce sync.Once
	file_eqxcustomers_sendotp_proto_rawDescData = file_eqxcustomers_sendotp_proto_rawDesc
)

func file_eqxcustomers_sendotp_proto_rawDescGZIP() []byte {
	file_eqxcustomers_sendotp_proto_rawDescOnce.Do(func() {
		file_eqxcustomers_sendotp_proto_rawDescData = protoimpl.X.CompressGZIP(file_eqxcustomers_sendotp_proto_rawDescData)
	})
	return file_eqxcustomers_sendotp_proto_rawDescData
}

var file_eqxcustomers_sendotp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eqxcustomers_sendotp_proto_goTypes = []interface{}{
	(*SendOtpRequest)(nil),  // 0: eqxcustomers.sendotp.sendOtpRequest
	(*SendOtpResponse)(nil), // 1: eqxcustomers.sendotp.sendOtpResponse
}
var file_eqxcustomers_sendotp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eqxcustomers_sendotp_proto_init() }
func file_eqxcustomers_sendotp_proto_init() {
	if File_eqxcustomers_sendotp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eqxcustomers_sendotp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOtpRequest); i {
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
		file_eqxcustomers_sendotp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOtpResponse); i {
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
			RawDescriptor: file_eqxcustomers_sendotp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eqxcustomers_sendotp_proto_goTypes,
		DependencyIndexes: file_eqxcustomers_sendotp_proto_depIdxs,
		MessageInfos:      file_eqxcustomers_sendotp_proto_msgTypes,
	}.Build()
	File_eqxcustomers_sendotp_proto = out.File
	file_eqxcustomers_sendotp_proto_rawDesc = nil
	file_eqxcustomers_sendotp_proto_goTypes = nil
	file_eqxcustomers_sendotp_proto_depIdxs = nil
}
