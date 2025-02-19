// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: lamf/getloandetails.proto

package lamf

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ReqId       string `protobuf:"bytes,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
	ClientRefNo string `protobuf:"bytes,3,opt,name=clientRefNo,proto3" json:"clientRefNo,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_getloandetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_getloandetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_lamf_getloandetails_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Request) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *Request) GetClientRefNo() string {
	if x != nil {
		return x.ClientRefNo
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode              *int32         `protobuf:"varint,1,opt,name=statusCode,proto3,oneof" json:"statusCode,omitempty"`
	TotatInvestedAmount     *float64       `protobuf:"fixed64,2,opt,name=totatInvestedAmount,proto3,oneof" json:"totatInvestedAmount,omitempty"`
	TotalCurrentMarketValue *float64       `protobuf:"fixed64,3,opt,name=totalCurrentMarketValue,proto3,oneof" json:"totalCurrentMarketValue,omitempty"`
	MaxLoanAmount           *float64       `protobuf:"fixed64,4,opt,name=maxLoanAmount,proto3,oneof" json:"maxLoanAmount,omitempty"`
	Data                    *Response_Data `protobuf:"bytes,5,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_getloandetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_getloandetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_lamf_getloandetails_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *Response) GetTotatInvestedAmount() float64 {
	if x != nil && x.TotatInvestedAmount != nil {
		return *x.TotatInvestedAmount
	}
	return 0
}

func (x *Response) GetTotalCurrentMarketValue() float64 {
	if x != nil && x.TotalCurrentMarketValue != nil {
		return *x.TotalCurrentMarketValue
	}
	return 0
}

func (x *Response) GetMaxLoanAmount() float64 {
	if x != nil && x.MaxLoanAmount != nil {
		return *x.MaxLoanAmount
	}
	return 0
}

func (x *Response) GetData() *Response_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Response_Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *int32  `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *Response_Errors) Reset() {
	*x = Response_Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_getloandetails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Errors) ProtoMessage() {}

func (x *Response_Errors) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_getloandetails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Errors.ProtoReflect.Descriptor instead.
func (*Response_Errors) Descriptor() ([]byte, []int) {
	return file_lamf_getloandetails_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_Errors) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *Response_Errors) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type Response_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId  *string            `protobuf:"bytes,1,opt,name=reqId,proto3,oneof" json:"reqId,omitempty"`
	Errors []*Response_Errors `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Response_Data) Reset() {
	*x = Response_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lamf_getloandetails_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Data) ProtoMessage() {}

func (x *Response_Data) ProtoReflect() protoreflect.Message {
	mi := &file_lamf_getloandetails_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Data.ProtoReflect.Descriptor instead.
func (*Response_Data) Descriptor() ([]byte, []int) {
	return file_lamf_getloandetails_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Response_Data) GetReqId() string {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return ""
}

func (x *Response_Data) GetErrors() []*Response_Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_lamf_getloandetails_proto protoreflect.FileDescriptor

var file_lamf_getloandetails_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x61, 0x6d, 0x66, 0x2f, 0x67, 0x65, 0x74, 0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6c, 0x61, 0x6d,
	0x66, 0x2e, 0x67, 0x65, 0x74, 0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x62, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x4e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x66, 0x4e, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x74, 0x49,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a,
	0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02,
	0x52, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x67, 0x65, 0x74,
	0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x1a, 0x55, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x17,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x69, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3c,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6c, 0x61, 0x6d, 0x66, 0x2e, 0x67, 0x65, 0x74, 0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x72, 0x65, 0x71, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x74, 0x49,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a, 0x0a,
	0x18, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x61,
	0x78, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lamf_getloandetails_proto_rawDescOnce sync.Once
	file_lamf_getloandetails_proto_rawDescData = file_lamf_getloandetails_proto_rawDesc
)

func file_lamf_getloandetails_proto_rawDescGZIP() []byte {
	file_lamf_getloandetails_proto_rawDescOnce.Do(func() {
		file_lamf_getloandetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_lamf_getloandetails_proto_rawDescData)
	})
	return file_lamf_getloandetails_proto_rawDescData
}

var file_lamf_getloandetails_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lamf_getloandetails_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: lamf.getloandetails.request
	(*Response)(nil),        // 1: lamf.getloandetails.response
	(*Response_Errors)(nil), // 2: lamf.getloandetails.response.Errors
	(*Response_Data)(nil),   // 3: lamf.getloandetails.response.Data
}
var file_lamf_getloandetails_proto_depIdxs = []int32{
	3, // 0: lamf.getloandetails.response.data:type_name -> lamf.getloandetails.response.Data
	2, // 1: lamf.getloandetails.response.Data.errors:type_name -> lamf.getloandetails.response.Errors
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lamf_getloandetails_proto_init() }
func file_lamf_getloandetails_proto_init() {
	if File_lamf_getloandetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lamf_getloandetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_lamf_getloandetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_lamf_getloandetails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Errors); i {
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
		file_lamf_getloandetails_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Data); i {
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
	file_lamf_getloandetails_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_lamf_getloandetails_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_lamf_getloandetails_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lamf_getloandetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lamf_getloandetails_proto_goTypes,
		DependencyIndexes: file_lamf_getloandetails_proto_depIdxs,
		MessageInfos:      file_lamf_getloandetails_proto_msgTypes,
	}.Build()
	File_lamf_getloandetails_proto = out.File
	file_lamf_getloandetails_proto_rawDesc = nil
	file_lamf_getloandetails_proto_goTypes = nil
	file_lamf_getloandetails_proto_depIdxs = nil
}
