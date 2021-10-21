// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/functions/grpc/pod-create.proto

package grpc

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

type CreatePodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info       *BaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	ActionID   *string   `protobuf:"bytes,2,opt,name=actionID,proto3,oneof" json:"actionID,omitempty"`
	InstanceID *string   `protobuf:"bytes,3,opt,name=InstanceID,proto3,oneof" json:"InstanceID,omitempty"`
	Step       *int64    `protobuf:"varint,4,opt,name=step,proto3,oneof" json:"step,omitempty"`
}

func (x *CreatePodRequest) Reset() {
	*x = CreatePodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_pod_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePodRequest) ProtoMessage() {}

func (x *CreatePodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_pod_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePodRequest.ProtoReflect.Descriptor instead.
func (*CreatePodRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_pod_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePodRequest) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CreatePodRequest) GetActionID() string {
	if x != nil && x.ActionID != nil {
		return *x.ActionID
	}
	return ""
}

func (x *CreatePodRequest) GetInstanceID() string {
	if x != nil && x.InstanceID != nil {
		return *x.InstanceID
	}
	return ""
}

func (x *CreatePodRequest) GetStep() int64 {
	if x != nil && x.Step != nil {
		return *x.Step
	}
	return 0
}

type CreatePodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       *string `protobuf:"bytes,1,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
	Hostname *string `protobuf:"bytes,2,opt,name=hostname,proto3,oneof" json:"hostname,omitempty"`
}

func (x *CreatePodResponse) Reset() {
	*x = CreatePodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_pod_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePodResponse) ProtoMessage() {}

func (x *CreatePodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_pod_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePodResponse.ProtoReflect.Descriptor instead.
func (*CreatePodResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_pod_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePodResponse) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *CreatePodResponse) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

var File_pkg_functions_grpc_pod_create_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_pod_create_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x64, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x22, 0x5d, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x70, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65,
	0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_pod_create_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_pod_create_proto_rawDescData = file_pkg_functions_grpc_pod_create_proto_rawDesc
)

func file_pkg_functions_grpc_pod_create_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_pod_create_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_pod_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_pod_create_proto_rawDescData)
	})
	return file_pkg_functions_grpc_pod_create_proto_rawDescData
}

var file_pkg_functions_grpc_pod_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_functions_grpc_pod_create_proto_goTypes = []interface{}{
	(*CreatePodRequest)(nil),  // 0: direktiv_functions.CreatePodRequest
	(*CreatePodResponse)(nil), // 1: direktiv_functions.CreatePodResponse
	(*BaseInfo)(nil),          // 2: direktiv_functions.BaseInfo
}
var file_pkg_functions_grpc_pod_create_proto_depIdxs = []int32{
	2, // 0: direktiv_functions.CreatePodRequest.info:type_name -> direktiv_functions.BaseInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_pod_create_proto_init() }
func file_pkg_functions_grpc_pod_create_proto_init() {
	if File_pkg_functions_grpc_pod_create_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_pod_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePodRequest); i {
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
		file_pkg_functions_grpc_pod_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePodResponse); i {
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
	file_pkg_functions_grpc_pod_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_pod_create_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_pod_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_pod_create_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_pod_create_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_pod_create_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_pod_create_proto = out.File
	file_pkg_functions_grpc_pod_create_proto_rawDesc = nil
	file_pkg_functions_grpc_pod_create_proto_goTypes = nil
	file_pkg_functions_grpc_pod_create_proto_depIdxs = nil
}
