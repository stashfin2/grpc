// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: ckyc/ckycdownload.proto

package ckyc

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

type CkycDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CkycReferenceId string `protobuf:"bytes,2,opt,name=ckyc_reference_id,json=ckycReferenceId,proto3" json:"ckyc_reference_id,omitempty"`
	Dob             string `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *CkycDownloadRequest) Reset() {
	*x = CkycDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ckyc_ckycdownload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CkycDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CkycDownloadRequest) ProtoMessage() {}

func (x *CkycDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ckyc_ckycdownload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CkycDownloadRequest.ProtoReflect.Descriptor instead.
func (*CkycDownloadRequest) Descriptor() ([]byte, []int) {
	return file_ckyc_ckycdownload_proto_rawDescGZIP(), []int{0}
}

func (x *CkycDownloadRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CkycDownloadRequest) GetCkycReferenceId() string {
	if x != nil {
		return x.CkycReferenceId
	}
	return ""
}

func (x *CkycDownloadRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

type CkycDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Xml    string `protobuf:"bytes,2,opt,name=xml,proto3" json:"xml,omitempty"`
}

func (x *CkycDownloadResponse) Reset() {
	*x = CkycDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ckyc_ckycdownload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CkycDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CkycDownloadResponse) ProtoMessage() {}

func (x *CkycDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ckyc_ckycdownload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CkycDownloadResponse.ProtoReflect.Descriptor instead.
func (*CkycDownloadResponse) Descriptor() ([]byte, []int) {
	return file_ckyc_ckycdownload_proto_rawDescGZIP(), []int{1}
}

func (x *CkycDownloadResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CkycDownloadResponse) GetXml() string {
	if x != nil {
		return x.Xml
	}
	return ""
}

var File_ckyc_ckycdownload_proto protoreflect.FileDescriptor

var file_ckyc_ckycdownload_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6b, 0x79, 0x63, 0x2f, 0x63, 0x6b, 0x79, 0x63, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6b, 0x79, 0x63, 0x2e,
	0x63, 0x6b, 0x79, 0x63, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x74, 0x0a, 0x13,
	0x43, 0x6b, 0x79, 0x63, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6b, 0x79, 0x63, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6b, 0x79, 0x63, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x6f, 0x62, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x6b, 0x79, 0x63, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x78, 0x6d, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ckyc_ckycdownload_proto_rawDescOnce sync.Once
	file_ckyc_ckycdownload_proto_rawDescData = file_ckyc_ckycdownload_proto_rawDesc
)

func file_ckyc_ckycdownload_proto_rawDescGZIP() []byte {
	file_ckyc_ckycdownload_proto_rawDescOnce.Do(func() {
		file_ckyc_ckycdownload_proto_rawDescData = protoimpl.X.CompressGZIP(file_ckyc_ckycdownload_proto_rawDescData)
	})
	return file_ckyc_ckycdownload_proto_rawDescData
}

var file_ckyc_ckycdownload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ckyc_ckycdownload_proto_goTypes = []interface{}{
	(*CkycDownloadRequest)(nil),  // 0: ckyc.ckycdownload.CkycDownloadRequest
	(*CkycDownloadResponse)(nil), // 1: ckyc.ckycdownload.CkycDownloadResponse
}
var file_ckyc_ckycdownload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ckyc_ckycdownload_proto_init() }
func file_ckyc_ckycdownload_proto_init() {
	if File_ckyc_ckycdownload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ckyc_ckycdownload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CkycDownloadRequest); i {
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
		file_ckyc_ckycdownload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CkycDownloadResponse); i {
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
			RawDescriptor: file_ckyc_ckycdownload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ckyc_ckycdownload_proto_goTypes,
		DependencyIndexes: file_ckyc_ckycdownload_proto_depIdxs,
		MessageInfos:      file_ckyc_ckycdownload_proto_msgTypes,
	}.Build()
	File_ckyc_ckycdownload_proto = out.File
	file_ckyc_ckycdownload_proto_rawDesc = nil
	file_ckyc_ckycdownload_proto_goTypes = nil
	file_ckyc_ckycdownload_proto_depIdxs = nil
}
