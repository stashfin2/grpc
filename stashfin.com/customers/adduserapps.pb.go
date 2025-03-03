// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: customers/adduserapps.proto

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

type UserApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName     string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppId       string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppCategory string `protobuf:"bytes,3,opt,name=app_category,json=appCategory,proto3" json:"app_category,omitempty"`
}

func (x *UserApp) Reset() {
	*x = UserApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_adduserapps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserApp) ProtoMessage() {}

func (x *UserApp) ProtoReflect() protoreflect.Message {
	mi := &file_customers_adduserapps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserApp.ProtoReflect.Descriptor instead.
func (*UserApp) Descriptor() ([]byte, []int) {
	return file_customers_adduserapps_proto_rawDescGZIP(), []int{0}
}

func (x *UserApp) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *UserApp) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UserApp) GetAppCategory() string {
	if x != nil {
		return x.AppCategory
	}
	return ""
}

type AddUserAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*UserApp `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *AddUserAppsRequest) Reset() {
	*x = AddUserAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_adduserapps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserAppsRequest) ProtoMessage() {}

func (x *AddUserAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_adduserapps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserAppsRequest.ProtoReflect.Descriptor instead.
func (*AddUserAppsRequest) Descriptor() ([]byte, []int) {
	return file_customers_adduserapps_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserAppsRequest) GetApps() []*UserApp {
	if x != nil {
		return x.Apps
	}
	return nil
}

type AddUserAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddUserAppsResponse) Reset() {
	*x = AddUserAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_adduserapps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserAppsResponse) ProtoMessage() {}

func (x *AddUserAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_adduserapps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserAppsResponse.ProtoReflect.Descriptor instead.
func (*AddUserAppsResponse) Descriptor() ([]byte, []int) {
	return file_customers_adduserapps_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserAppsResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_customers_adduserapps_proto protoreflect.FileDescriptor

var file_customers_adduserapps_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x75, 0x73, 0x65, 0x72,
	0x61, 0x70, 0x70, 0x73, 0x22, 0x5e, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0x48, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x61, 0x70,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x2d,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x77, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x64, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x12,
	0x29, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_adduserapps_proto_rawDescOnce sync.Once
	file_customers_adduserapps_proto_rawDescData = file_customers_adduserapps_proto_rawDesc
)

func file_customers_adduserapps_proto_rawDescGZIP() []byte {
	file_customers_adduserapps_proto_rawDescOnce.Do(func() {
		file_customers_adduserapps_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_adduserapps_proto_rawDescData)
	})
	return file_customers_adduserapps_proto_rawDescData
}

var file_customers_adduserapps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customers_adduserapps_proto_goTypes = []interface{}{
	(*UserApp)(nil),             // 0: customers.adduserapps.UserApp
	(*AddUserAppsRequest)(nil),  // 1: customers.adduserapps.AddUserAppsRequest
	(*AddUserAppsResponse)(nil), // 2: customers.adduserapps.AddUserAppsResponse
}
var file_customers_adduserapps_proto_depIdxs = []int32{
	0, // 0: customers.adduserapps.AddUserAppsRequest.apps:type_name -> customers.adduserapps.UserApp
	1, // 1: customers.adduserapps.UserAppsService.AddUserApps:input_type -> customers.adduserapps.AddUserAppsRequest
	2, // 2: customers.adduserapps.UserAppsService.AddUserApps:output_type -> customers.adduserapps.AddUserAppsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customers_adduserapps_proto_init() }
func file_customers_adduserapps_proto_init() {
	if File_customers_adduserapps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_adduserapps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserApp); i {
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
		file_customers_adduserapps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserAppsRequest); i {
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
		file_customers_adduserapps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserAppsResponse); i {
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
			RawDescriptor: file_customers_adduserapps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customers_adduserapps_proto_goTypes,
		DependencyIndexes: file_customers_adduserapps_proto_depIdxs,
		MessageInfos:      file_customers_adduserapps_proto_msgTypes,
	}.Build()
	File_customers_adduserapps_proto = out.File
	file_customers_adduserapps_proto_rawDesc = nil
	file_customers_adduserapps_proto_goTypes = nil
	file_customers_adduserapps_proto_depIdxs = nil
}
