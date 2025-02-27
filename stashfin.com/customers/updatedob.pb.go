// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/updatedob.proto

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

type UpdateDOBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dob string `protobuf:"bytes,1,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *UpdateDOBRequest) Reset() {
	*x = UpdateDOBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_updatedob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDOBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDOBRequest) ProtoMessage() {}

func (x *UpdateDOBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatedob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDOBRequest.ProtoReflect.Descriptor instead.
func (*UpdateDOBRequest) Descriptor() ([]byte, []int) {
	return file_customers_updatedob_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDOBRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

type UpdateDOBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateDOBResponse) Reset() {
	*x = UpdateDOBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_updatedob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDOBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDOBResponse) ProtoMessage() {}

func (x *UpdateDOBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatedob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDOBResponse.ProtoReflect.Descriptor instead.
func (*UpdateDOBResponse) Descriptor() ([]byte, []int) {
	return file_customers_updatedob_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDOBResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_customers_updatedob_proto protoreflect.FileDescriptor

var file_customers_updatedob_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x62,
	0x22, 0x24, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x4f, 0x42, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x22, 0x2b, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x4f, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_updatedob_proto_rawDescOnce sync.Once
	file_customers_updatedob_proto_rawDescData = file_customers_updatedob_proto_rawDesc
)

func file_customers_updatedob_proto_rawDescGZIP() []byte {
	file_customers_updatedob_proto_rawDescOnce.Do(func() {
		file_customers_updatedob_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_updatedob_proto_rawDescData)
	})
	return file_customers_updatedob_proto_rawDescData
}

var file_customers_updatedob_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_updatedob_proto_goTypes = []interface{}{
	(*UpdateDOBRequest)(nil),  // 0: customers.updatedob.updateDOBRequest
	(*UpdateDOBResponse)(nil), // 1: customers.updatedob.updateDOBResponse
}
var file_customers_updatedob_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_updatedob_proto_init() }
func file_customers_updatedob_proto_init() {
	if File_customers_updatedob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_updatedob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDOBRequest); i {
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
		file_customers_updatedob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDOBResponse); i {
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
			RawDescriptor: file_customers_updatedob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_updatedob_proto_goTypes,
		DependencyIndexes: file_customers_updatedob_proto_depIdxs,
		MessageInfos:      file_customers_updatedob_proto_msgTypes,
	}.Build()
	File_customers_updatedob_proto = out.File
	file_customers_updatedob_proto_rawDesc = nil
	file_customers_updatedob_proto_goTypes = nil
	file_customers_updatedob_proto_depIdxs = nil
}
