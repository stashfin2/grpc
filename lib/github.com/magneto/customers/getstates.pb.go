// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: customers/getstates.proto

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

type GetAllStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllStatesRequest) Reset() {
	*x = GetAllStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStatesRequest) ProtoMessage() {}

func (x *GetAllStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStatesRequest.ProtoReflect.Descriptor instead.
func (*GetAllStatesRequest) Descriptor() ([]byte, []int) {
	return file_customers_getstates_proto_rawDescGZIP(), []int{0}
}

type GetAllStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *GetAllStatesResponse) Reset() {
	*x = GetAllStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStatesResponse) ProtoMessage() {}

func (x *GetAllStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStatesResponse.ProtoReflect.Descriptor instead.
func (*GetAllStatesResponse) Descriptor() ([]byte, []int) {
	return file_customers_getstates_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllStatesResponse) GetStates() []*State {
	if x != nil {
		return x.States
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_customers_getstates_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *State) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_customers_getstates_proto protoreflect.FileDescriptor

var file_customers_getstates_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x15, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_getstates_proto_rawDescOnce sync.Once
	file_customers_getstates_proto_rawDescData = file_customers_getstates_proto_rawDesc
)

func file_customers_getstates_proto_rawDescGZIP() []byte {
	file_customers_getstates_proto_rawDescOnce.Do(func() {
		file_customers_getstates_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_getstates_proto_rawDescData)
	})
	return file_customers_getstates_proto_rawDescData
}

var file_customers_getstates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_getstates_proto_goTypes = []interface{}{
	(*GetAllStatesRequest)(nil),  // 0: customers.getstates.getAllStatesRequest
	(*GetAllStatesResponse)(nil), // 1: customers.getstates.getAllStatesResponse
	(*State)(nil),                // 2: customers.getstates.State
}
var file_customers_getstates_proto_depIdxs = []int32{
	2, // 0: customers.getstates.getAllStatesResponse.states:type_name -> customers.getstates.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_getstates_proto_init() }
func file_customers_getstates_proto_init() {
	if File_customers_getstates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_getstates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStatesRequest); i {
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
		file_customers_getstates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStatesResponse); i {
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
		file_customers_getstates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_customers_getstates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_getstates_proto_goTypes,
		DependencyIndexes: file_customers_getstates_proto_depIdxs,
		MessageInfos:      file_customers_getstates_proto_msgTypes,
	}.Build()
	File_customers_getstates_proto = out.File
	file_customers_getstates_proto_rawDesc = nil
	file_customers_getstates_proto_goTypes = nil
	file_customers_getstates_proto_depIdxs = nil
}
