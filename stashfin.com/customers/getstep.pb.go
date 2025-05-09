// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/getstep.proto

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

type GetstepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerCode *string `protobuf:"bytes,1,opt,name=banner_code,json=bannerCode,proto3,oneof" json:"banner_code,omitempty"`
}

func (x *GetstepRequest) Reset() {
	*x = GetstepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetstepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetstepRequest) ProtoMessage() {}

func (x *GetstepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetstepRequest.ProtoReflect.Descriptor instead.
func (*GetstepRequest) Descriptor() ([]byte, []int) {
	return file_customers_getstep_proto_rawDescGZIP(), []int{0}
}

func (x *GetstepRequest) GetBannerCode() string {
	if x != nil && x.BannerCode != nil {
		return *x.BannerCode
	}
	return ""
}

type GetstepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route      string                   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Type       string                   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ButtonText string                   `protobuf:"bytes,3,opt,name=button_text,json=buttonText,proto3" json:"button_text,omitempty"`
	Data       []*GetstepResponse_Field `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetstepResponse) Reset() {
	*x = GetstepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetstepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetstepResponse) ProtoMessage() {}

func (x *GetstepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetstepResponse.ProtoReflect.Descriptor instead.
func (*GetstepResponse) Descriptor() ([]byte, []int) {
	return file_customers_getstep_proto_rawDescGZIP(), []int{1}
}

func (x *GetstepResponse) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *GetstepResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetstepResponse) GetButtonText() string {
	if x != nil {
		return x.ButtonText
	}
	return ""
}

func (x *GetstepResponse) GetData() []*GetstepResponse_Field {
	if x != nil {
		return x.Data
	}
	return nil
}

type StepValidationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StepValidationData) Reset() {
	*x = StepValidationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepValidationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepValidationData) ProtoMessage() {}

func (x *StepValidationData) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepValidationData.ProtoReflect.Descriptor instead.
func (*StepValidationData) Descriptor() ([]byte, []int) {
	return file_customers_getstep_proto_rawDescGZIP(), []int{2}
}

func (x *StepValidationData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StepValidationData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StepValidationData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetstepResponse_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label              string                `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Key                string                `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	InputType          string                `protobuf:"bytes,3,opt,name=input_type,json=inputType,proto3" json:"input_type,omitempty"`
	DataType           string                `protobuf:"bytes,4,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Icon               string                `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	IsPopup            bool                  `protobuf:"varint,6,opt,name=is_popup,json=isPopup,proto3" json:"is_popup,omitempty"`
	KeyboardType       string                `protobuf:"bytes,7,opt,name=keyboard_type,json=keyboardType,proto3" json:"keyboard_type,omitempty"`
	TextCapitalization string                `protobuf:"bytes,8,opt,name=text_capitalization,json=textCapitalization,proto3" json:"text_capitalization,omitempty"`
	Validations        []*StepValidationData `protobuf:"bytes,9,rep,name=validations,proto3" json:"validations,omitempty"`
}

func (x *GetstepResponse_Field) Reset() {
	*x = GetstepResponse_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_getstep_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetstepResponse_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetstepResponse_Field) ProtoMessage() {}

func (x *GetstepResponse_Field) ProtoReflect() protoreflect.Message {
	mi := &file_customers_getstep_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetstepResponse_Field.ProtoReflect.Descriptor instead.
func (*GetstepResponse_Field) Descriptor() ([]byte, []int) {
	return file_customers_getstep_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetstepResponse_Field) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *GetstepResponse_Field) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetstepResponse_Field) GetInputType() string {
	if x != nil {
		return x.InputType
	}
	return ""
}

func (x *GetstepResponse_Field) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *GetstepResponse_Field) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *GetstepResponse_Field) GetIsPopup() bool {
	if x != nil {
		return x.IsPopup
	}
	return false
}

func (x *GetstepResponse_Field) GetKeyboardType() string {
	if x != nil {
		return x.KeyboardType
	}
	return ""
}

func (x *GetstepResponse_Field) GetTextCapitalization() string {
	if x != nil {
		return x.TextCapitalization
	}
	return ""
}

func (x *GetstepResponse_Field) GetValidations() []*StepValidationData {
	if x != nil {
		return x.Validations
	}
	return nil
}

var File_customers_getstep_proto protoreflect.FileDescriptor

var file_customers_getstep_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73,
	0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x74, 0x65, 0x70, 0x22, 0x46, 0x0a, 0x0e,
	0x67, 0x65, 0x74, 0x73, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0xd6, 0x03, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x73, 0x74, 0x65, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65,
	0x74, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0xb9, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x70, 0x75, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x6f, 0x70, 0x75, 0x70, 0x12, 0x23,
	0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x74, 0x65, 0x78, 0x74, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56, 0x0a,
	0x12, 0x53, 0x74, 0x65, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_getstep_proto_rawDescOnce sync.Once
	file_customers_getstep_proto_rawDescData = file_customers_getstep_proto_rawDesc
)

func file_customers_getstep_proto_rawDescGZIP() []byte {
	file_customers_getstep_proto_rawDescOnce.Do(func() {
		file_customers_getstep_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_getstep_proto_rawDescData)
	})
	return file_customers_getstep_proto_rawDescData
}

var file_customers_getstep_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_customers_getstep_proto_goTypes = []interface{}{
	(*GetstepRequest)(nil),        // 0: customers.getstep.getstepRequest
	(*GetstepResponse)(nil),       // 1: customers.getstep.getstepResponse
	(*StepValidationData)(nil),    // 2: customers.getstep.StepValidationData
	(*GetstepResponse_Field)(nil), // 3: customers.getstep.getstepResponse.Field
}
var file_customers_getstep_proto_depIdxs = []int32{
	3, // 0: customers.getstep.getstepResponse.data:type_name -> customers.getstep.getstepResponse.Field
	2, // 1: customers.getstep.getstepResponse.Field.validations:type_name -> customers.getstep.StepValidationData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_customers_getstep_proto_init() }
func file_customers_getstep_proto_init() {
	if File_customers_getstep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_getstep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetstepRequest); i {
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
		file_customers_getstep_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetstepResponse); i {
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
		file_customers_getstep_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepValidationData); i {
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
		file_customers_getstep_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetstepResponse_Field); i {
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
	file_customers_getstep_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customers_getstep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_getstep_proto_goTypes,
		DependencyIndexes: file_customers_getstep_proto_depIdxs,
		MessageInfos:      file_customers_getstep_proto_msgTypes,
	}.Build()
	File_customers_getstep_proto = out.File
	file_customers_getstep_proto_rawDesc = nil
	file_customers_getstep_proto_goTypes = nil
	file_customers_getstep_proto_depIdxs = nil
}
