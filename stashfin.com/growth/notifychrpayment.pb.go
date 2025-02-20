// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: growth/notifychrpayment.proto

package growth

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

type Notifychrrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	OrderId    string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId int32   `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Mode       *string `protobuf:"bytes,4,opt,name=mode,proto3,oneof" json:"mode,omitempty"`
	MetaData   string  `protobuf:"bytes,5,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	ClientId   string  `protobuf:"bytes,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *Notifychrrequest) Reset() {
	*x = Notifychrrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_growth_notifychrpayment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifychrrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifychrrequest) ProtoMessage() {}

func (x *Notifychrrequest) ProtoReflect() protoreflect.Message {
	mi := &file_growth_notifychrpayment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifychrrequest.ProtoReflect.Descriptor instead.
func (*Notifychrrequest) Descriptor() ([]byte, []int) {
	return file_growth_notifychrpayment_proto_rawDescGZIP(), []int{0}
}

func (x *Notifychrrequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Notifychrrequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Notifychrrequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Notifychrrequest) GetMode() string {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return ""
}

func (x *Notifychrrequest) GetMetaData() string {
	if x != nil {
		return x.MetaData
	}
	return ""
}

func (x *Notifychrrequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Notifychrresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Notifychrresponse) Reset() {
	*x = Notifychrresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_growth_notifychrpayment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifychrresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifychrresponse) ProtoMessage() {}

func (x *Notifychrresponse) ProtoReflect() protoreflect.Message {
	mi := &file_growth_notifychrpayment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifychrresponse.ProtoReflect.Descriptor instead.
func (*Notifychrresponse) Descriptor() ([]byte, []int) {
	return file_growth_notifychrpayment_proto_rawDescGZIP(), []int{1}
}

func (x *Notifychrresponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Notifychrresponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Notifychrresponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_growth_notifychrpayment_proto protoreflect.FileDescriptor

var file_growth_notifychrpayment_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63,
	0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63, 0x68,
	0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x65, 0x0a,
	0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_growth_notifychrpayment_proto_rawDescOnce sync.Once
	file_growth_notifychrpayment_proto_rawDescData = file_growth_notifychrpayment_proto_rawDesc
)

func file_growth_notifychrpayment_proto_rawDescGZIP() []byte {
	file_growth_notifychrpayment_proto_rawDescOnce.Do(func() {
		file_growth_notifychrpayment_proto_rawDescData = protoimpl.X.CompressGZIP(file_growth_notifychrpayment_proto_rawDescData)
	})
	return file_growth_notifychrpayment_proto_rawDescData
}

var file_growth_notifychrpayment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_growth_notifychrpayment_proto_goTypes = []interface{}{
	(*Notifychrrequest)(nil),  // 0: growth.notifychrpayment.notifychrrequest
	(*Notifychrresponse)(nil), // 1: growth.notifychrpayment.notifychrresponse
}
var file_growth_notifychrpayment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_growth_notifychrpayment_proto_init() }
func file_growth_notifychrpayment_proto_init() {
	if File_growth_notifychrpayment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_growth_notifychrpayment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifychrrequest); i {
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
		file_growth_notifychrpayment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifychrresponse); i {
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
	file_growth_notifychrpayment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_growth_notifychrpayment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_growth_notifychrpayment_proto_goTypes,
		DependencyIndexes: file_growth_notifychrpayment_proto_depIdxs,
		MessageInfos:      file_growth_notifychrpayment_proto_msgTypes,
	}.Build()
	File_growth_notifychrpayment_proto = out.File
	file_growth_notifychrpayment_proto_rawDesc = nil
	file_growth_notifychrpayment_proto_goTypes = nil
	file_growth_notifychrpayment_proto_depIdxs = nil
}
