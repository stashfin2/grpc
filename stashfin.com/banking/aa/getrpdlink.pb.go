// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: banking/aa/getrpdlink.proto

package aa

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

type CreateRpdLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *CreateRpdLinkRequest) Reset() {
	*x = CreateRpdLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_getrpdlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRpdLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRpdLinkRequest) ProtoMessage() {}

func (x *CreateRpdLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_getrpdlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRpdLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateRpdLinkRequest) Descriptor() ([]byte, []int) {
	return file_banking_aa_getrpdlink_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRpdLinkRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type CreateRpdLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortUrl  string `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	TraceId   string `protobuf:"bytes,4,opt,name=traceId,proto3" json:"traceId,omitempty"`
	UpiBillId string `protobuf:"bytes,5,opt,name=upiBillId,proto3" json:"upiBillId,omitempty"`
	UpiLink   string `protobuf:"bytes,6,opt,name=upiLink,proto3" json:"upiLink,omitempty"`
	ValidUpto string `protobuf:"bytes,7,opt,name=validUpto,proto3" json:"validUpto,omitempty"`
}

func (x *CreateRpdLinkResponse) Reset() {
	*x = CreateRpdLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_aa_getrpdlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRpdLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRpdLinkResponse) ProtoMessage() {}

func (x *CreateRpdLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_aa_getrpdlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRpdLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateRpdLinkResponse) Descriptor() ([]byte, []int) {
	return file_banking_aa_getrpdlink_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRpdLinkResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetUpiBillId() string {
	if x != nil {
		return x.UpiBillId
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetUpiLink() string {
	if x != nil {
		return x.UpiLink
	}
	return ""
}

func (x *CreateRpdLinkResponse) GetValidUpto() string {
	if x != nil {
		return x.ValidUpto
	}
	return ""
}

var File_banking_aa_getrpdlink_proto protoreflect.FileDescriptor

var file_banking_aa_getrpdlink_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x61, 0x2f, 0x67, 0x65, 0x74,
	0x72, 0x70, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x72, 0x70, 0x64,
	0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x37, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x70,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0xcb, 0x01,
	0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x70, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x69, 0x42, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x69, 0x42, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x69, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x69, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x70, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x70, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_banking_aa_getrpdlink_proto_rawDescOnce sync.Once
	file_banking_aa_getrpdlink_proto_rawDescData = file_banking_aa_getrpdlink_proto_rawDesc
)

func file_banking_aa_getrpdlink_proto_rawDescGZIP() []byte {
	file_banking_aa_getrpdlink_proto_rawDescOnce.Do(func() {
		file_banking_aa_getrpdlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_banking_aa_getrpdlink_proto_rawDescData)
	})
	return file_banking_aa_getrpdlink_proto_rawDescData
}

var file_banking_aa_getrpdlink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_banking_aa_getrpdlink_proto_goTypes = []interface{}{
	(*CreateRpdLinkRequest)(nil),  // 0: banking.aa.getrpdlink.createRpdLinkRequest
	(*CreateRpdLinkResponse)(nil), // 1: banking.aa.getrpdlink.createRpdLinkResponse
}
var file_banking_aa_getrpdlink_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_banking_aa_getrpdlink_proto_init() }
func file_banking_aa_getrpdlink_proto_init() {
	if File_banking_aa_getrpdlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banking_aa_getrpdlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRpdLinkRequest); i {
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
		file_banking_aa_getrpdlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRpdLinkResponse); i {
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
			RawDescriptor: file_banking_aa_getrpdlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banking_aa_getrpdlink_proto_goTypes,
		DependencyIndexes: file_banking_aa_getrpdlink_proto_depIdxs,
		MessageInfos:      file_banking_aa_getrpdlink_proto_msgTypes,
	}.Build()
	File_banking_aa_getrpdlink_proto = out.File
	file_banking_aa_getrpdlink_proto_rawDesc = nil
	file_banking_aa_getrpdlink_proto_goTypes = nil
	file_banking_aa_getrpdlink_proto_depIdxs = nil
}
