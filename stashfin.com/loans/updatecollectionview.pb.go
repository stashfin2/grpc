// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: loans/updatecollectionview.proto

package loans

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

type UpdateCollectionViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         *string `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	BounceReason   *string `protobuf:"bytes,2,opt,name=bounce_reason,json=bounceReason,proto3,oneof" json:"bounce_reason,omitempty"`
	PtpDate        *string `protobuf:"bytes,3,opt,name=ptp_date,json=ptpDate,proto3,oneof" json:"ptp_date,omitempty"`
	AgentName      *string `protobuf:"bytes,4,opt,name=agent_name,json=agentName,proto3,oneof" json:"agent_name,omitempty"`
	FosAgent       *string `protobuf:"bytes,5,opt,name=fos_agent,json=fosAgent,proto3,oneof" json:"fos_agent,omitempty"`
	ClLeadAssignId *int32  `protobuf:"varint,6,opt,name=cl_lead_assign_id,json=clLeadAssignId,proto3,oneof" json:"cl_lead_assign_id,omitempty"`
	Level          *int32  `protobuf:"varint,7,opt,name=level,proto3,oneof" json:"level,omitempty"`
	AssignedLevel  *int32  `protobuf:"varint,8,opt,name=assigned_level,json=assignedLevel,proto3,oneof" json:"assigned_level,omitempty"`
	CustomerId     int32   `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *UpdateCollectionViewRequest) Reset() {
	*x = UpdateCollectionViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_updatecollectionview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionViewRequest) ProtoMessage() {}

func (x *UpdateCollectionViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_updatecollectionview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionViewRequest.ProtoReflect.Descriptor instead.
func (*UpdateCollectionViewRequest) Descriptor() ([]byte, []int) {
	return file_loans_updatecollectionview_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCollectionViewRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *UpdateCollectionViewRequest) GetBounceReason() string {
	if x != nil && x.BounceReason != nil {
		return *x.BounceReason
	}
	return ""
}

func (x *UpdateCollectionViewRequest) GetPtpDate() string {
	if x != nil && x.PtpDate != nil {
		return *x.PtpDate
	}
	return ""
}

func (x *UpdateCollectionViewRequest) GetAgentName() string {
	if x != nil && x.AgentName != nil {
		return *x.AgentName
	}
	return ""
}

func (x *UpdateCollectionViewRequest) GetFosAgent() string {
	if x != nil && x.FosAgent != nil {
		return *x.FosAgent
	}
	return ""
}

func (x *UpdateCollectionViewRequest) GetClLeadAssignId() int32 {
	if x != nil && x.ClLeadAssignId != nil {
		return *x.ClLeadAssignId
	}
	return 0
}

func (x *UpdateCollectionViewRequest) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *UpdateCollectionViewRequest) GetAssignedLevel() int32 {
	if x != nil && x.AssignedLevel != nil {
		return *x.AssignedLevel
	}
	return 0
}

func (x *UpdateCollectionViewRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type UpdateCollectionViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id     int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateCollectionViewResponse) Reset() {
	*x = UpdateCollectionViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_updatecollectionview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionViewResponse) ProtoMessage() {}

func (x *UpdateCollectionViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_updatecollectionview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionViewResponse.ProtoReflect.Descriptor instead.
func (*UpdateCollectionViewResponse) Descriptor() ([]byte, []int) {
	return file_loans_updatecollectionview_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCollectionViewResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateCollectionViewResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_loans_updatecollectionview_proto protoreflect.FileDescriptor

var file_loans_updatecollectionview_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x69, 0x65, 0x77, 0x22, 0xdc,
	0x03, 0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x62,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x70, 0x74, 0x70, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x70, 0x74, 0x70, 0x44, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x6f, 0x73,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08,
	0x66, 0x6f, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x11, 0x63,
	0x6c, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x0e, 0x63, 0x6c, 0x4c, 0x65, 0x61, 0x64,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x07,
	0x52, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x74, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x66, 0x6f, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63,
	0x6c, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x46, 0x0a,
	0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_updatecollectionview_proto_rawDescOnce sync.Once
	file_loans_updatecollectionview_proto_rawDescData = file_loans_updatecollectionview_proto_rawDesc
)

func file_loans_updatecollectionview_proto_rawDescGZIP() []byte {
	file_loans_updatecollectionview_proto_rawDescOnce.Do(func() {
		file_loans_updatecollectionview_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_updatecollectionview_proto_rawDescData)
	})
	return file_loans_updatecollectionview_proto_rawDescData
}

var file_loans_updatecollectionview_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loans_updatecollectionview_proto_goTypes = []interface{}{
	(*UpdateCollectionViewRequest)(nil),  // 0: loans.updatecollectionview.updateCollectionViewRequest
	(*UpdateCollectionViewResponse)(nil), // 1: loans.updatecollectionview.updateCollectionViewResponse
}
var file_loans_updatecollectionview_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loans_updatecollectionview_proto_init() }
func file_loans_updatecollectionview_proto_init() {
	if File_loans_updatecollectionview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_updatecollectionview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionViewRequest); i {
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
		file_loans_updatecollectionview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionViewResponse); i {
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
	file_loans_updatecollectionview_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loans_updatecollectionview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_updatecollectionview_proto_goTypes,
		DependencyIndexes: file_loans_updatecollectionview_proto_depIdxs,
		MessageInfos:      file_loans_updatecollectionview_proto_msgTypes,
	}.Build()
	File_loans_updatecollectionview_proto = out.File
	file_loans_updatecollectionview_proto_rawDesc = nil
	file_loans_updatecollectionview_proto_goTypes = nil
	file_loans_updatecollectionview_proto_depIdxs = nil
}
