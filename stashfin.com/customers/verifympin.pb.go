// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/verifympin.proto

package customers

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

type VerifyMpinRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Mpin          int32                  `protobuf:"varint,2,opt,name=mpin,proto3" json:"mpin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyMpinRequest) Reset() {
	*x = VerifyMpinRequest{}
	mi := &file_customers_verifympin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyMpinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyMpinRequest) ProtoMessage() {}

func (x *VerifyMpinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_verifympin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyMpinRequest.ProtoReflect.Descriptor instead.
func (*VerifyMpinRequest) Descriptor() ([]byte, []int) {
	return file_customers_verifympin_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyMpinRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *VerifyMpinRequest) GetMpin() int32 {
	if x != nil {
		return x.Mpin
	}
	return 0
}

type VerifyMpinResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Jwttoken      string                 `protobuf:"bytes,2,opt,name=jwttoken,proto3" json:"jwttoken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyMpinResponse) Reset() {
	*x = VerifyMpinResponse{}
	mi := &file_customers_verifympin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyMpinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyMpinResponse) ProtoMessage() {}

func (x *VerifyMpinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_verifympin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyMpinResponse.ProtoReflect.Descriptor instead.
func (*VerifyMpinResponse) Descriptor() ([]byte, []int) {
	return file_customers_verifympin_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyMpinResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *VerifyMpinResponse) GetJwttoken() string {
	if x != nil {
		return x.Jwttoken
	}
	return ""
}

var File_customers_verifympin_proto protoreflect.FileDescriptor

var file_customers_verifympin_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x6d, 0x70, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x6d, 0x70,
	0x69, 0x6e, 0x22, 0x3f, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x70, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d,
	0x70, 0x69, 0x6e, 0x22, 0x48, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x70, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x77, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_verifympin_proto_rawDescOnce sync.Once
	file_customers_verifympin_proto_rawDescData []byte
)

func file_customers_verifympin_proto_rawDescGZIP() []byte {
	file_customers_verifympin_proto_rawDescOnce.Do(func() {
		file_customers_verifympin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_verifympin_proto_rawDesc), len(file_customers_verifympin_proto_rawDesc)))
	})
	return file_customers_verifympin_proto_rawDescData
}

var file_customers_verifympin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_verifympin_proto_goTypes = []any{
	(*VerifyMpinRequest)(nil),  // 0: customers.verifympin.verifyMpinRequest
	(*VerifyMpinResponse)(nil), // 1: customers.verifympin.verifyMpinResponse
}
var file_customers_verifympin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_verifympin_proto_init() }
func file_customers_verifympin_proto_init() {
	if File_customers_verifympin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_verifympin_proto_rawDesc), len(file_customers_verifympin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_verifympin_proto_goTypes,
		DependencyIndexes: file_customers_verifympin_proto_depIdxs,
		MessageInfos:      file_customers_verifympin_proto_msgTypes,
	}.Build()
	File_customers_verifympin_proto = out.File
	file_customers_verifympin_proto_goTypes = nil
	file_customers_verifympin_proto_depIdxs = nil
}
