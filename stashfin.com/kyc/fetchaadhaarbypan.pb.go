// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: kyc/fetchaadhaarbypan.proto

package kyc

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

type FetchAadhaarByPanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanNumber string `protobuf:"bytes,1,opt,name=pan_number,json=panNumber,proto3" json:"pan_number,omitempty"`
}

func (x *FetchAadhaarByPanRequest) Reset() {
	*x = FetchAadhaarByPanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAadhaarByPanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAadhaarByPanRequest) ProtoMessage() {}

func (x *FetchAadhaarByPanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAadhaarByPanRequest.ProtoReflect.Descriptor instead.
func (*FetchAadhaarByPanRequest) Descriptor() ([]byte, []int) {
	return file_kyc_fetchaadhaarbypan_proto_rawDescGZIP(), []int{0}
}

func (x *FetchAadhaarByPanRequest) GetPanNumber() string {
	if x != nil {
		return x.PanNumber
	}
	return ""
}

type FetchAadhaarByPanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId     string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status        int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Data          *Data  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp     uint64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Path          string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FetchAadhaarByPanResponse) Reset() {
	*x = FetchAadhaarByPanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAadhaarByPanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAadhaarByPanResponse) ProtoMessage() {}

func (x *FetchAadhaarByPanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAadhaarByPanResponse.ProtoReflect.Descriptor instead.
func (*FetchAadhaarByPanResponse) Descriptor() ([]byte, []int) {
	return file_kyc_fetchaadhaarbypan_proto_rawDescGZIP(), []int{1}
}

func (x *FetchAadhaarByPanResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FetchAadhaarByPanResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *FetchAadhaarByPanResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FetchAadhaarByPanResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FetchAadhaarByPanResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *FetchAadhaarByPanResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	PanData *PanData `protobuf:"bytes,3,opt,name=pan_data,json=panData,proto3,oneof" json:"pan_data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_kyc_fetchaadhaarbypan_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Data) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Data) GetPanData() *PanData {
	if x != nil {
		return x.PanData
	}
	return nil
}

type PanData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentType        string `protobuf:"bytes,1,opt,name=document_type,json=documentType,proto3" json:"document_type,omitempty"`
	DocumentId          string `protobuf:"bytes,2,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	MaskedAadhaarNumber string `protobuf:"bytes,3,opt,name=masked_aadhaar_number,json=maskedAadhaarNumber,proto3" json:"masked_aadhaar_number,omitempty"`
}

func (x *PanData) Reset() {
	*x = PanData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanData) ProtoMessage() {}

func (x *PanData) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_fetchaadhaarbypan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanData.ProtoReflect.Descriptor instead.
func (*PanData) Descriptor() ([]byte, []int) {
	return file_kyc_fetchaadhaarbypan_proto_rawDescGZIP(), []int{3}
}

func (x *PanData) GetDocumentType() string {
	if x != nil {
		return x.DocumentType
	}
	return ""
}

func (x *PanData) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *PanData) GetMaskedAadhaarNumber() string {
	if x != nil {
		return x.MaskedAadhaarNumber
	}
	return ""
}

var File_kyc_fetchaadhaarbypan_proto protoreflect.FileDescriptor

var file_kyc_fetchaadhaarbypan_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6b, 0x79, 0x63, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x61, 0x61, 0x64, 0x68, 0x61,
	0x61, 0x72, 0x62, 0x79, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b,
	0x79, 0x63, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x61, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x62,
	0x79, 0x70, 0x61, 0x6e, 0x22, 0x39, 0x0a, 0x18, 0x66, 0x65, 0x74, 0x63, 0x68, 0x41, 0x61, 0x64,
	0x68, 0x61, 0x61, 0x72, 0x42, 0x79, 0x50, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0xdc, 0x01, 0x0a, 0x19, 0x66, 0x65, 0x74, 0x63, 0x68, 0x41, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72,
	0x42, 0x79, 0x50, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x79, 0x63, 0x2e,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x61, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x62, 0x79, 0x70, 0x61,
	0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x81,
	0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x61, 0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x62, 0x79, 0x70, 0x61, 0x6e, 0x2e,
	0x50, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x61,
	0x61, 0x64, 0x68, 0x61, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x41, 0x61, 0x64, 0x68, 0x61,
	0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kyc_fetchaadhaarbypan_proto_rawDescOnce sync.Once
	file_kyc_fetchaadhaarbypan_proto_rawDescData = file_kyc_fetchaadhaarbypan_proto_rawDesc
)

func file_kyc_fetchaadhaarbypan_proto_rawDescGZIP() []byte {
	file_kyc_fetchaadhaarbypan_proto_rawDescOnce.Do(func() {
		file_kyc_fetchaadhaarbypan_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyc_fetchaadhaarbypan_proto_rawDescData)
	})
	return file_kyc_fetchaadhaarbypan_proto_rawDescData
}

var file_kyc_fetchaadhaarbypan_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kyc_fetchaadhaarbypan_proto_goTypes = []interface{}{
	(*FetchAadhaarByPanRequest)(nil),  // 0: kyc.fetchaadhaarbypan.fetchAadhaarByPanRequest
	(*FetchAadhaarByPanResponse)(nil), // 1: kyc.fetchaadhaarbypan.fetchAadhaarByPanResponse
	(*Data)(nil),                      // 2: kyc.fetchaadhaarbypan.Data
	(*PanData)(nil),                   // 3: kyc.fetchaadhaarbypan.PanData
}
var file_kyc_fetchaadhaarbypan_proto_depIdxs = []int32{
	2, // 0: kyc.fetchaadhaarbypan.fetchAadhaarByPanResponse.data:type_name -> kyc.fetchaadhaarbypan.Data
	3, // 1: kyc.fetchaadhaarbypan.Data.pan_data:type_name -> kyc.fetchaadhaarbypan.PanData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kyc_fetchaadhaarbypan_proto_init() }
func file_kyc_fetchaadhaarbypan_proto_init() {
	if File_kyc_fetchaadhaarbypan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyc_fetchaadhaarbypan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAadhaarByPanRequest); i {
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
		file_kyc_fetchaadhaarbypan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAadhaarByPanResponse); i {
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
		file_kyc_fetchaadhaarbypan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_kyc_fetchaadhaarbypan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanData); i {
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
	file_kyc_fetchaadhaarbypan_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kyc_fetchaadhaarbypan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kyc_fetchaadhaarbypan_proto_goTypes,
		DependencyIndexes: file_kyc_fetchaadhaarbypan_proto_depIdxs,
		MessageInfos:      file_kyc_fetchaadhaarbypan_proto_msgTypes,
	}.Build()
	File_kyc_fetchaadhaarbypan_proto = out.File
	file_kyc_fetchaadhaarbypan_proto_rawDesc = nil
	file_kyc_fetchaadhaarbypan_proto_goTypes = nil
	file_kyc_fetchaadhaarbypan_proto_depIdxs = nil
}
