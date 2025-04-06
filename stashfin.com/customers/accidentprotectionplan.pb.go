// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: customers/accidentprotectionplan.proto

package customers

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

type Benefit struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Icon          string                 `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Benefit) Reset() {
	*x = Benefit{}
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Benefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benefit) ProtoMessage() {}

func (x *Benefit) ProtoReflect() protoreflect.Message {
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benefit.ProtoReflect.Descriptor instead.
func (*Benefit) Descriptor() ([]byte, []int) {
	return file_customers_accidentprotectionplan_proto_rawDescGZIP(), []int{0}
}

func (x *Benefit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Benefit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Benefit) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type GetPocketPersonalAccidentProtectionPlanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPocketPersonalAccidentProtectionPlanRequest) Reset() {
	*x = GetPocketPersonalAccidentProtectionPlanRequest{}
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPocketPersonalAccidentProtectionPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPocketPersonalAccidentProtectionPlanRequest) ProtoMessage() {}

func (x *GetPocketPersonalAccidentProtectionPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPocketPersonalAccidentProtectionPlanRequest.ProtoReflect.Descriptor instead.
func (*GetPocketPersonalAccidentProtectionPlanRequest) Descriptor() ([]byte, []int) {
	return file_customers_accidentprotectionplan_proto_rawDescGZIP(), []int{1}
}

type GetPocketPersonalAccidentProtectionPlanResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Title            string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Benefits         []*Benefit             `protobuf:"bytes,2,rep,name=benefits,proto3" json:"benefits,omitempty"`
	TermsDescription string                 `protobuf:"bytes,3,opt,name=terms_description,json=termsDescription,proto3" json:"terms_description,omitempty"`
	TermsLink        string                 `protobuf:"bytes,4,opt,name=terms_link,json=termsLink,proto3" json:"terms_link,omitempty"`
	ButtonName       string                 `protobuf:"bytes,5,opt,name=button_name,json=buttonName,proto3" json:"button_name,omitempty"`
	Status           string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	PlanName         string                 `protobuf:"bytes,7,opt,name=plan_name,json=planName,proto3" json:"plan_name,omitempty"`
	PlanDetails      *PlanDetails           `protobuf:"bytes,8,opt,name=plan_details,json=planDetails,proto3" json:"plan_details,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) Reset() {
	*x = GetPocketPersonalAccidentProtectionPlanResponse{}
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPocketPersonalAccidentProtectionPlanResponse) ProtoMessage() {}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPocketPersonalAccidentProtectionPlanResponse.ProtoReflect.Descriptor instead.
func (*GetPocketPersonalAccidentProtectionPlanResponse) Descriptor() ([]byte, []int) {
	return file_customers_accidentprotectionplan_proto_rawDescGZIP(), []int{2}
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetBenefits() []*Benefit {
	if x != nil {
		return x.Benefits
	}
	return nil
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetTermsDescription() string {
	if x != nil {
		return x.TermsDescription
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetTermsLink() string {
	if x != nil {
		return x.TermsLink
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetButtonName() string {
	if x != nil {
		return x.ButtonName
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetPlanName() string {
	if x != nil {
		return x.PlanName
	}
	return ""
}

func (x *GetPocketPersonalAccidentProtectionPlanResponse) GetPlanDetails() *PlanDetails {
	if x != nil {
		return x.PlanDetails
	}
	return nil
}

type PlanDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PlanId        int32                  `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Amount        float32                `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	MetaData      string                 `protobuf:"bytes,4,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanDetails) Reset() {
	*x = PlanDetails{}
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanDetails) ProtoMessage() {}

func (x *PlanDetails) ProtoReflect() protoreflect.Message {
	mi := &file_customers_accidentprotectionplan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanDetails.ProtoReflect.Descriptor instead.
func (*PlanDetails) Descriptor() ([]byte, []int) {
	return file_customers_accidentprotectionplan_proto_rawDescGZIP(), []int{3}
}

func (x *PlanDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PlanDetails) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *PlanDetails) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PlanDetails) GetMetaData() string {
	if x != nil {
		return x.MetaData
	}
	return ""
}

var File_customers_accidentprotectionplan_proto protoreflect.FileDescriptor

var file_customers_accidentprotectionplan_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x55, 0x0a, 0x07, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x22, 0x30, 0x0a, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x82, 0x03, 0x0a, 0x2f, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x45, 0x0a,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x08, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x6f, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_customers_accidentprotectionplan_proto_rawDescOnce sync.Once
	file_customers_accidentprotectionplan_proto_rawDescData []byte
)

func file_customers_accidentprotectionplan_proto_rawDescGZIP() []byte {
	file_customers_accidentprotectionplan_proto_rawDescOnce.Do(func() {
		file_customers_accidentprotectionplan_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customers_accidentprotectionplan_proto_rawDesc), len(file_customers_accidentprotectionplan_proto_rawDesc)))
	})
	return file_customers_accidentprotectionplan_proto_rawDescData
}

var file_customers_accidentprotectionplan_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_customers_accidentprotectionplan_proto_goTypes = []any{
	(*Benefit)(nil), // 0: customers.accidentprotectionplan.Benefit
	(*GetPocketPersonalAccidentProtectionPlanRequest)(nil),  // 1: customers.accidentprotectionplan.getPocketPersonalAccidentProtectionPlanRequest
	(*GetPocketPersonalAccidentProtectionPlanResponse)(nil), // 2: customers.accidentprotectionplan.getPocketPersonalAccidentProtectionPlanResponse
	(*PlanDetails)(nil), // 3: customers.accidentprotectionplan.PlanDetails
}
var file_customers_accidentprotectionplan_proto_depIdxs = []int32{
	0, // 0: customers.accidentprotectionplan.getPocketPersonalAccidentProtectionPlanResponse.benefits:type_name -> customers.accidentprotectionplan.Benefit
	3, // 1: customers.accidentprotectionplan.getPocketPersonalAccidentProtectionPlanResponse.plan_details:type_name -> customers.accidentprotectionplan.PlanDetails
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_customers_accidentprotectionplan_proto_init() }
func file_customers_accidentprotectionplan_proto_init() {
	if File_customers_accidentprotectionplan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customers_accidentprotectionplan_proto_rawDesc), len(file_customers_accidentprotectionplan_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_accidentprotectionplan_proto_goTypes,
		DependencyIndexes: file_customers_accidentprotectionplan_proto_depIdxs,
		MessageInfos:      file_customers_accidentprotectionplan_proto_msgTypes,
	}.Build()
	File_customers_accidentprotectionplan_proto = out.File
	file_customers_accidentprotectionplan_proto_goTypes = nil
	file_customers_accidentprotectionplan_proto_depIdxs = nil
}
