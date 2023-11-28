// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/goremote.proto

package proto

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

type ScreenshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Victim string `protobuf:"bytes,1,opt,name=victim,proto3" json:"victim,omitempty"`
}

func (x *ScreenshotRequest) Reset() {
	*x = ScreenshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goremote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenshotRequest) ProtoMessage() {}

func (x *ScreenshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goremote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenshotRequest.ProtoReflect.Descriptor instead.
func (*ScreenshotRequest) Descriptor() ([]byte, []int) {
	return file_proto_goremote_proto_rawDescGZIP(), []int{0}
}

func (x *ScreenshotRequest) GetVictim() string {
	if x != nil {
		return x.Victim
	}
	return ""
}

type ScreenshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ScreenshotResponse) Reset() {
	*x = ScreenshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goremote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenshotResponse) ProtoMessage() {}

func (x *ScreenshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goremote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenshotResponse.ProtoReflect.Descriptor instead.
func (*ScreenshotResponse) Descriptor() ([]byte, []int) {
	return file_proto_goremote_proto_rawDescGZIP(), []int{1}
}

func (x *ScreenshotResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *ScreenshotResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_goremote_proto protoreflect.FileDescriptor

var file_proto_goremote_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x11, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6d, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x59, 0x0a, 0x0d,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x48, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_goremote_proto_rawDescOnce sync.Once
	file_proto_goremote_proto_rawDescData = file_proto_goremote_proto_rawDesc
)

func file_proto_goremote_proto_rawDescGZIP() []byte {
	file_proto_goremote_proto_rawDescOnce.Do(func() {
		file_proto_goremote_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_goremote_proto_rawDescData)
	})
	return file_proto_goremote_proto_rawDescData
}

var file_proto_goremote_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_goremote_proto_goTypes = []interface{}{
	(*ScreenshotRequest)(nil),  // 0: proto.ScreenshotRequest
	(*ScreenshotResponse)(nil), // 1: proto.ScreenshotResponse
}
var file_proto_goremote_proto_depIdxs = []int32{
	0, // 0: proto.RemoteControl.GetScreenshot:input_type -> proto.ScreenshotRequest
	1, // 1: proto.RemoteControl.GetScreenshot:output_type -> proto.ScreenshotResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_goremote_proto_init() }
func file_proto_goremote_proto_init() {
	if File_proto_goremote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_goremote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenshotRequest); i {
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
		file_proto_goremote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenshotResponse); i {
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
			RawDescriptor: file_proto_goremote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_goremote_proto_goTypes,
		DependencyIndexes: file_proto_goremote_proto_depIdxs,
		MessageInfos:      file_proto_goremote_proto_msgTypes,
	}.Build()
	File_proto_goremote_proto = out.File
	file_proto_goremote_proto_rawDesc = nil
	file_proto_goremote_proto_goTypes = nil
	file_proto_goremote_proto_depIdxs = nil
}
