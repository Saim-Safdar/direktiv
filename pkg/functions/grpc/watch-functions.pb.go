// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/functions/grpc/watch-functions.proto

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

type Traffic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RevisionName *string `protobuf:"bytes,1,opt,name=revisionName,proto3,oneof" json:"revisionName,omitempty"`
	Traffic      *int64  `protobuf:"varint,2,opt,name=traffic,proto3,oneof" json:"traffic,omitempty"`
	Generation   *int64  `protobuf:"varint,3,opt,name=generation,proto3,oneof" json:"generation,omitempty"`
}

func (x *Traffic) Reset() {
	*x = Traffic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traffic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traffic) ProtoMessage() {}

func (x *Traffic) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traffic.ProtoReflect.Descriptor instead.
func (*Traffic) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_functions_proto_rawDescGZIP(), []int{0}
}

func (x *Traffic) GetRevisionName() string {
	if x != nil && x.RevisionName != nil {
		return *x.RevisionName
	}
	return ""
}

func (x *Traffic) GetTraffic() int64 {
	if x != nil && x.Traffic != nil {
		return *x.Traffic
	}
	return 0
}

func (x *Traffic) GetGeneration() int64 {
	if x != nil && x.Generation != nil {
		return *x.Generation
	}
	return 0
}

type WatchFunctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WatchFunctionsRequest) Reset() {
	*x = WatchFunctionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchFunctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchFunctionsRequest) ProtoMessage() {}

func (x *WatchFunctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchFunctionsRequest.ProtoReflect.Descriptor instead.
func (*WatchFunctionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_functions_proto_rawDescGZIP(), []int{1}
}

func (x *WatchFunctionsRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type WatchFunctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    *string        `protobuf:"bytes,1,opt,name=event,proto3,oneof" json:"event,omitempty"`
	Function *FunctionsInfo `protobuf:"bytes,2,opt,name=function,proto3,oneof" json:"function,omitempty"`
	Traffic  []*Traffic     `protobuf:"bytes,3,rep,name=traffic,proto3" json:"traffic,omitempty"`
}

func (x *WatchFunctionsResponse) Reset() {
	*x = WatchFunctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchFunctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchFunctionsResponse) ProtoMessage() {}

func (x *WatchFunctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_watch_functions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchFunctionsResponse.ProtoReflect.Descriptor instead.
func (*WatchFunctionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_watch_functions_proto_rawDescGZIP(), []int{2}
}

func (x *WatchFunctionsResponse) GetEvent() string {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return ""
}

func (x *WatchFunctionsResponse) GetFunction() *FunctionsInfo {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *WatchFunctionsResponse) GetTraffic() []*Traffic {
	if x != nil {
		return x.Traffic
	}
	return nil
}

var File_pkg_functions_grpc_watch_functions_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_watch_functions_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x07,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0c, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x01, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xa7, 0x01, 0x0a, 0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9, 0x01, 0x0a, 0x16, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_watch_functions_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_watch_functions_proto_rawDescData = file_pkg_functions_grpc_watch_functions_proto_rawDesc
)

func file_pkg_functions_grpc_watch_functions_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_watch_functions_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_watch_functions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_watch_functions_proto_rawDescData)
	})
	return file_pkg_functions_grpc_watch_functions_proto_rawDescData
}

var file_pkg_functions_grpc_watch_functions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_functions_grpc_watch_functions_proto_goTypes = []interface{}{
	(*Traffic)(nil),                // 0: grpc.Traffic
	(*WatchFunctionsRequest)(nil),  // 1: grpc.WatchFunctionsRequest
	(*WatchFunctionsResponse)(nil), // 2: grpc.WatchFunctionsResponse
	nil,                            // 3: grpc.WatchFunctionsRequest.AnnotationsEntry
	(*FunctionsInfo)(nil),          // 4: grpc.FunctionsInfo
}
var file_pkg_functions_grpc_watch_functions_proto_depIdxs = []int32{
	3, // 0: grpc.WatchFunctionsRequest.annotations:type_name -> grpc.WatchFunctionsRequest.AnnotationsEntry
	4, // 1: grpc.WatchFunctionsResponse.function:type_name -> grpc.FunctionsInfo
	0, // 2: grpc.WatchFunctionsResponse.traffic:type_name -> grpc.Traffic
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_watch_functions_proto_init() }
func file_pkg_functions_grpc_watch_functions_proto_init() {
	if File_pkg_functions_grpc_watch_functions_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_watch_functions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traffic); i {
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
		file_pkg_functions_grpc_watch_functions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchFunctionsRequest); i {
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
		file_pkg_functions_grpc_watch_functions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchFunctionsResponse); i {
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
	file_pkg_functions_grpc_watch_functions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_watch_functions_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_watch_functions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_watch_functions_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_watch_functions_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_watch_functions_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_watch_functions_proto = out.File
	file_pkg_functions_grpc_watch_functions_proto_rawDesc = nil
	file_pkg_functions_grpc_watch_functions_proto_goTypes = nil
	file_pkg_functions_grpc_watch_functions_proto_depIdxs = nil
}
