// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/aa/getbanklist.proto

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

type GetBankListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBankListRequest) Reset() {
	*x = GetBankListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_aa_getbanklist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBankListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBankListRequest) ProtoMessage() {}

func (x *GetBankListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_aa_getbanklist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBankListRequest.ProtoReflect.Descriptor instead.
func (*GetBankListRequest) Descriptor() ([]byte, []int) {
	return file_customers_aa_getbanklist_proto_rawDescGZIP(), []int{0}
}

type GetBankListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fip           []*BankData `protobuf:"bytes,1,rep,name=fip,proto3" json:"fip,omitempty"`
	NetbankingUrl string      `protobuf:"bytes,2,opt,name=netbanking_url,json=netbankingUrl,proto3" json:"netbanking_url,omitempty"`
}

func (x *GetBankListResponse) Reset() {
	*x = GetBankListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_aa_getbanklist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBankListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBankListResponse) ProtoMessage() {}

func (x *GetBankListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_aa_getbanklist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBankListResponse.ProtoReflect.Descriptor instead.
func (*GetBankListResponse) Descriptor() ([]byte, []int) {
	return file_customers_aa_getbanklist_proto_rawDescGZIP(), []int{1}
}

func (x *GetBankListResponse) GetFip() []*BankData {
	if x != nil {
		return x.Fip
	}
	return nil
}

func (x *GetBankListResponse) GetNetbankingUrl() string {
	if x != nil {
		return x.NetbankingUrl
	}
	return ""
}

type BankData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NetBankingEnabled bool   `protobuf:"varint,3,opt,name=netBankingEnabled,proto3" json:"netBankingEnabled,omitempty"`
	PdfJourneyEnabled bool   `protobuf:"varint,4,opt,name=pdfJourneyEnabled,proto3" json:"pdfJourneyEnabled,omitempty"`
}

func (x *BankData) Reset() {
	*x = BankData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_aa_getbanklist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankData) ProtoMessage() {}

func (x *BankData) ProtoReflect() protoreflect.Message {
	mi := &file_customers_aa_getbanklist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankData.ProtoReflect.Descriptor instead.
func (*BankData) Descriptor() ([]byte, []int) {
	return file_customers_aa_getbanklist_proto_rawDescGZIP(), []int{2}
}

func (x *BankData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankData) GetNetBankingEnabled() bool {
	if x != nil {
		return x.NetBankingEnabled
	}
	return false
}

func (x *BankData) GetPdfJourneyEnabled() bool {
	if x != nil {
		return x.PdfJourneyEnabled
	}
	return false
}

var File_customers_aa_getbanklist_proto protoreflect.FileDescriptor

var file_customers_aa_getbanklist_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x61, 0x2f, 0x67,
	0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x61, 0x2e, 0x67,
	0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x67, 0x65,
	0x74, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x72, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x66, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2e, 0x61, 0x61, 0x2e, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03, 0x66, 0x69, 0x70, 0x12, 0x25, 0x0a,
	0x0e, 0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x55, 0x72, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x6e, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x64, 0x66, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x70, 0x64, 0x66, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_aa_getbanklist_proto_rawDescOnce sync.Once
	file_customers_aa_getbanklist_proto_rawDescData = file_customers_aa_getbanklist_proto_rawDesc
)

func file_customers_aa_getbanklist_proto_rawDescGZIP() []byte {
	file_customers_aa_getbanklist_proto_rawDescOnce.Do(func() {
		file_customers_aa_getbanklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_aa_getbanklist_proto_rawDescData)
	})
	return file_customers_aa_getbanklist_proto_rawDescData
}

var file_customers_aa_getbanklist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_aa_getbanklist_proto_goTypes = []interface{}{
	(*GetBankListRequest)(nil),  // 0: customers.aa.getbanklist.getBankListRequest
	(*GetBankListResponse)(nil), // 1: customers.aa.getbanklist.getBankListResponse
	(*BankData)(nil),            // 2: customers.aa.getbanklist.BankData
}
var file_customers_aa_getbanklist_proto_depIdxs = []int32{
	2, // 0: customers.aa.getbanklist.getBankListResponse.fip:type_name -> customers.aa.getbanklist.BankData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_aa_getbanklist_proto_init() }
func file_customers_aa_getbanklist_proto_init() {
	if File_customers_aa_getbanklist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_aa_getbanklist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBankListRequest); i {
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
		file_customers_aa_getbanklist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBankListResponse); i {
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
		file_customers_aa_getbanklist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankData); i {
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
			RawDescriptor: file_customers_aa_getbanklist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_aa_getbanklist_proto_goTypes,
		DependencyIndexes: file_customers_aa_getbanklist_proto_depIdxs,
		MessageInfos:      file_customers_aa_getbanklist_proto_msgTypes,
	}.Build()
	File_customers_aa_getbanklist_proto = out.File
	file_customers_aa_getbanklist_proto_rawDesc = nil
	file_customers_aa_getbanklist_proto_goTypes = nil
	file_customers_aa_getbanklist_proto_depIdxs = nil
}
