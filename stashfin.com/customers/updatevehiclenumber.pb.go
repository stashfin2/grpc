// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/updatevehiclenumber.proto

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

type UpdateVehicleNumberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RcNumber      string                 `protobuf:"bytes,1,opt,name=rcNumber,proto3" json:"rcNumber,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVehicleNumberRequest) Reset() {
	*x = UpdateVehicleNumberRequest{}
	mi := &file_customers_updatevehiclenumber_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVehicleNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVehicleNumberRequest) ProtoMessage() {}

func (x *UpdateVehicleNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatevehiclenumber_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVehicleNumberRequest.ProtoReflect.Descriptor instead.
func (*UpdateVehicleNumberRequest) Descriptor() ([]byte, []int) {
	return file_customers_updatevehiclenumber_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateVehicleNumberRequest) GetRcNumber() string {
	if x != nil {
		return x.RcNumber
	}
	return ""
}

type UpdateVehicleNumberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVehicleNumberResponse) Reset() {
	*x = UpdateVehicleNumberResponse{}
	mi := &file_customers_updatevehiclenumber_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVehicleNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVehicleNumberResponse) ProtoMessage() {}

func (x *UpdateVehicleNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_updatevehiclenumber_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVehicleNumberResponse.ProtoReflect.Descriptor instead.
func (*UpdateVehicleNumberResponse) Descriptor() ([]byte, []int) {
	return file_customers_updatevehiclenumber_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateVehicleNumberResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_customers_updatevehiclenumber_proto protoreflect.FileDescriptor

var file_customers_updatevehiclenumber_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x1a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x35,
	0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_customers_updatevehiclenumber_proto_rawDescOnce sync.Once
	file_customers_updatevehiclenumber_proto_rawDescData []byte
)

func file_customers_updatevehiclenumber_proto_rawDescGZIP() []byte {
	file_customers_updatevehiclenumber_proto_rawDescOnce.Do(func() {
		file_customers_updatevehiclenumber_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_updatevehiclenumber_proto_rawDesc), len(file_customers_updatevehiclenumber_proto_rawDesc)))
	})
	return file_customers_updatevehiclenumber_proto_rawDescData
}

var file_customers_updatevehiclenumber_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_updatevehiclenumber_proto_goTypes = []any{
	(*UpdateVehicleNumberRequest)(nil),  // 0: customers.updatevehiclenumber.updateVehicleNumberRequest
	(*UpdateVehicleNumberResponse)(nil), // 1: customers.updatevehiclenumber.updateVehicleNumberResponse
}
var file_customers_updatevehiclenumber_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_updatevehiclenumber_proto_init() }
func file_customers_updatevehiclenumber_proto_init() {
	if File_customers_updatevehiclenumber_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_updatevehiclenumber_proto_rawDesc), len(file_customers_updatevehiclenumber_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_updatevehiclenumber_proto_goTypes,
		DependencyIndexes: file_customers_updatevehiclenumber_proto_depIdxs,
		MessageInfos:      file_customers_updatevehiclenumber_proto_msgTypes,
	}.Build()
	File_customers_updatevehiclenumber_proto = out.File
	file_customers_updatevehiclenumber_proto_goTypes = nil
	file_customers_updatevehiclenumber_proto_depIdxs = nil
}
