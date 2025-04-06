// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: documents/getagreement.proto

package documents

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

type Documentrequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LoanId        int32                  `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Documentrequest) Reset() {
	*x = Documentrequest{}
	mi := &file_documents_getagreement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Documentrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documentrequest) ProtoMessage() {}

func (x *Documentrequest) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getagreement_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documentrequest.ProtoReflect.Descriptor instead.
func (*Documentrequest) Descriptor() ([]byte, []int) {
	return file_documents_getagreement_proto_rawDescGZIP(), []int{0}
}

func (x *Documentrequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

type Agreementdata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgreementUrl  string                 `protobuf:"bytes,1,opt,name=agreement_url,json=agreementUrl,proto3" json:"agreement_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Agreementdata) Reset() {
	*x = Agreementdata{}
	mi := &file_documents_getagreement_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Agreementdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agreementdata) ProtoMessage() {}

func (x *Agreementdata) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getagreement_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agreementdata.ProtoReflect.Descriptor instead.
func (*Agreementdata) Descriptor() ([]byte, []int) {
	return file_documents_getagreement_proto_rawDescGZIP(), []int{1}
}

func (x *Agreementdata) GetAgreementUrl() string {
	if x != nil {
		return x.AgreementUrl
	}
	return ""
}

type Documentresponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Data          *Agreementdata         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Documentresponse) Reset() {
	*x = Documentresponse{}
	mi := &file_documents_getagreement_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Documentresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documentresponse) ProtoMessage() {}

func (x *Documentresponse) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getagreement_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documentresponse.ProtoReflect.Descriptor instead.
func (*Documentresponse) Descriptor() ([]byte, []int) {
	return file_documents_getagreement_proto_rawDescGZIP(), []int{2}
}

func (x *Documentresponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Documentresponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Documentresponse) GetData() *Agreementdata {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Documentresponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_documents_getagreement_proto protoreflect.FileDescriptor

var file_documents_getagreement_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e,
	0x49, 0x64, 0x22, 0x34, 0x0a, 0x0d, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x67, 0x65, 0x74, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x72,
	0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_documents_getagreement_proto_rawDescOnce sync.Once
	file_documents_getagreement_proto_rawDescData []byte
)

func file_documents_getagreement_proto_rawDescGZIP() []byte {
	file_documents_getagreement_proto_rawDescOnce.Do(func() {
		file_documents_getagreement_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_documents_getagreement_proto_rawDesc), len(file_documents_getagreement_proto_rawDesc)))
	})
	return file_documents_getagreement_proto_rawDescData
}

var file_documents_getagreement_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_documents_getagreement_proto_goTypes = []any{
	(*Documentrequest)(nil),  // 0: documents.getagreement.documentrequest
	(*Agreementdata)(nil),    // 1: documents.getagreement.agreementdata
	(*Documentresponse)(nil), // 2: documents.getagreement.documentresponse
}
var file_documents_getagreement_proto_depIdxs = []int32{
	1, // 0: documents.getagreement.documentresponse.data:type_name -> documents.getagreement.agreementdata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_documents_getagreement_proto_init() }
func file_documents_getagreement_proto_init() {
	if File_documents_getagreement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_documents_getagreement_proto_rawDesc), len(file_documents_getagreement_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_documents_getagreement_proto_goTypes,
		DependencyIndexes: file_documents_getagreement_proto_depIdxs,
		MessageInfos:      file_documents_getagreement_proto_msgTypes,
	}.Build()
	File_documents_getagreement_proto = out.File
	file_documents_getagreement_proto_goTypes = nil
	file_documents_getagreement_proto_depIdxs = nil
}
