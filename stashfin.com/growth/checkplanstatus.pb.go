// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: growth/checkplanstatus.proto

package growth

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

type Chrplansstatusrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chrplansstatusrequest) Reset() {
	*x = Chrplansstatusrequest{}
	mi := &file_growth_checkplanstatus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chrplansstatusrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chrplansstatusrequest) ProtoMessage() {}

func (x *Chrplansstatusrequest) ProtoReflect() protoreflect.Message {
	mi := &file_growth_checkplanstatus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chrplansstatusrequest.ProtoReflect.Descriptor instead.
func (*Chrplansstatusrequest) Descriptor() ([]byte, []int) {
	return file_growth_checkplanstatus_proto_rawDescGZIP(), []int{0}
}

type Planstatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlanStatus    bool                   `protobuf:"varint,1,opt,name=planStatus,proto3" json:"planStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Planstatus) Reset() {
	*x = Planstatus{}
	mi := &file_growth_checkplanstatus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Planstatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Planstatus) ProtoMessage() {}

func (x *Planstatus) ProtoReflect() protoreflect.Message {
	mi := &file_growth_checkplanstatus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Planstatus.ProtoReflect.Descriptor instead.
func (*Planstatus) Descriptor() ([]byte, []int) {
	return file_growth_checkplanstatus_proto_rawDescGZIP(), []int{1}
}

func (x *Planstatus) GetPlanStatus() bool {
	if x != nil {
		return x.PlanStatus
	}
	return false
}

type Chrplanstatusresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Data          *Planstatus            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chrplanstatusresponse) Reset() {
	*x = Chrplanstatusresponse{}
	mi := &file_growth_checkplanstatus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chrplanstatusresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chrplanstatusresponse) ProtoMessage() {}

func (x *Chrplanstatusresponse) ProtoReflect() protoreflect.Message {
	mi := &file_growth_checkplanstatus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chrplanstatusresponse.ProtoReflect.Descriptor instead.
func (*Chrplanstatusresponse) Descriptor() ([]byte, []int) {
	return file_growth_checkplanstatus_proto_rawDescGZIP(), []int{2}
}

func (x *Chrplanstatusresponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Chrplanstatusresponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Chrplanstatusresponse) GetData() *Planstatus {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Chrplanstatusresponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_growth_checkplanstatus_proto protoreflect.FileDescriptor

var file_growth_checkplanstatus_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c,
	0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c, 0x61, 0x6e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2c, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1, 0x01,
	0x0a, 0x15, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c, 0x61, 0x6e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_growth_checkplanstatus_proto_rawDescOnce sync.Once
	file_growth_checkplanstatus_proto_rawDescData []byte
)

func file_growth_checkplanstatus_proto_rawDescGZIP() []byte {
	file_growth_checkplanstatus_proto_rawDescOnce.Do(func() {
		file_growth_checkplanstatus_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_growth_checkplanstatus_proto_rawDesc), len(file_growth_checkplanstatus_proto_rawDesc)))
	})
	return file_growth_checkplanstatus_proto_rawDescData
}

var file_growth_checkplanstatus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_growth_checkplanstatus_proto_goTypes = []any{
	(*Chrplansstatusrequest)(nil), // 0: growth.checkplanstatus.chrplansstatusrequest
	(*Planstatus)(nil),            // 1: growth.checkplanstatus.planstatus
	(*Chrplanstatusresponse)(nil), // 2: growth.checkplanstatus.chrplanstatusresponse
}
var file_growth_checkplanstatus_proto_depIdxs = []int32{
	1, // 0: growth.checkplanstatus.chrplanstatusresponse.data:type_name -> growth.checkplanstatus.planstatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_growth_checkplanstatus_proto_init() }
func file_growth_checkplanstatus_proto_init() {
	if File_growth_checkplanstatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_growth_checkplanstatus_proto_rawDesc), len(file_growth_checkplanstatus_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_growth_checkplanstatus_proto_goTypes,
		DependencyIndexes: file_growth_checkplanstatus_proto_depIdxs,
		MessageInfos:      file_growth_checkplanstatus_proto_msgTypes,
	}.Build()
	File_growth_checkplanstatus_proto = out.File
	file_growth_checkplanstatus_proto_goTypes = nil
	file_growth_checkplanstatus_proto_depIdxs = nil
}
