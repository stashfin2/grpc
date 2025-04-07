// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: kyc/facematch.proto

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

type FaceMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfieImage string `protobuf:"bytes,1,opt,name=selfie_image,json=selfieImage,proto3" json:"selfie_image,omitempty"`
	KycImage    string `protobuf:"bytes,2,opt,name=kyc_image,json=kycImage,proto3" json:"kyc_image,omitempty"`
	CustomerId  int32  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	TxnId       string `protobuf:"bytes,4,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
}

func (x *FaceMatchRequest) Reset() {
	*x = FaceMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_facematch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceMatchRequest) ProtoMessage() {}

func (x *FaceMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_facematch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceMatchRequest.ProtoReflect.Descriptor instead.
func (*FaceMatchRequest) Descriptor() ([]byte, []int) {
	return file_kyc_facematch_proto_rawDescGZIP(), []int{0}
}

func (x *FaceMatchRequest) GetSelfieImage() string {
	if x != nil {
		return x.SelfieImage
	}
	return ""
}

func (x *FaceMatchRequest) GetKycImage() string {
	if x != nil {
		return x.KycImage
	}
	return ""
}

func (x *FaceMatchRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FaceMatchRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

type FaceMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf         string  `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	Match        string  `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	MatchScore   int32   `protobuf:"varint,3,opt,name=match_score,json=matchScore,proto3" json:"match_score,omitempty"`
	ToBeReviewed string  `protobuf:"bytes,4,opt,name=to_be_reviewed,json=toBeReviewed,proto3" json:"to_be_reviewed,omitempty"`
	CustomerId   int32   `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	TxnId        string  `protobuf:"bytes,6,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	Status       string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode   int32   `protobuf:"varint,8,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessage *string `protobuf:"bytes,9,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *FaceMatchResponse) Reset() {
	*x = FaceMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_facematch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceMatchResponse) ProtoMessage() {}

func (x *FaceMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_facematch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceMatchResponse.ProtoReflect.Descriptor instead.
func (*FaceMatchResponse) Descriptor() ([]byte, []int) {
	return file_kyc_facematch_proto_rawDescGZIP(), []int{1}
}

func (x *FaceMatchResponse) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

func (x *FaceMatchResponse) GetMatch() string {
	if x != nil {
		return x.Match
	}
	return ""
}

func (x *FaceMatchResponse) GetMatchScore() int32 {
	if x != nil {
		return x.MatchScore
	}
	return 0
}

func (x *FaceMatchResponse) GetToBeReviewed() string {
	if x != nil {
		return x.ToBeReviewed
	}
	return ""
}

func (x *FaceMatchResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FaceMatchResponse) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *FaceMatchResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FaceMatchResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FaceMatchResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

var File_kyc_facematch_proto protoreflect.FileDescriptor

var file_kyc_facematch_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x79, 0x63, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b, 0x79, 0x63, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6c,
	0x66, 0x69, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6b, 0x79, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x79, 0x63, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49,
	0x64, 0x22, 0xb1, 0x02, 0x0a, 0x11, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6f, 0x5f, 0x62, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x42, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kyc_facematch_proto_rawDescOnce sync.Once
	file_kyc_facematch_proto_rawDescData = file_kyc_facematch_proto_rawDesc
)

func file_kyc_facematch_proto_rawDescGZIP() []byte {
	file_kyc_facematch_proto_rawDescOnce.Do(func() {
		file_kyc_facematch_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyc_facematch_proto_rawDescData)
	})
	return file_kyc_facematch_proto_rawDescData
}

var file_kyc_facematch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kyc_facematch_proto_goTypes = []interface{}{
	(*FaceMatchRequest)(nil),  // 0: kyc.facematch.faceMatchRequest
	(*FaceMatchResponse)(nil), // 1: kyc.facematch.faceMatchResponse
}
var file_kyc_facematch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kyc_facematch_proto_init() }
func file_kyc_facematch_proto_init() {
	if File_kyc_facematch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyc_facematch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceMatchRequest); i {
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
		file_kyc_facematch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceMatchResponse); i {
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
	file_kyc_facematch_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kyc_facematch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kyc_facematch_proto_goTypes,
		DependencyIndexes: file_kyc_facematch_proto_depIdxs,
		MessageInfos:      file_kyc_facematch_proto_msgTypes,
	}.Build()
	File_kyc_facematch_proto = out.File
	file_kyc_facematch_proto_rawDesc = nil
	file_kyc_facematch_proto_goTypes = nil
	file_kyc_facematch_proto_depIdxs = nil
}
