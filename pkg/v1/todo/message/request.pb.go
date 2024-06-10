// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.0
// source: v1/todo/message/request.proto

package message

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	TaskName string `protobuf:"bytes,1,opt,name=taskName,proto3" json:"taskName,omitempty" validate:"required"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *CreateRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_v1_todo_message_request_proto protoreflect.FileDescriptor

var file_v1_todo_message_request_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0x92, 0x41, 0x0e, 0x4a, 0x0c, 0x22, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92,
	0x41, 0x08, 0x4a, 0x06, 0x22, 0x54, 0x4f, 0x44, 0x4f, 0x22, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x77, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2d, 0x6d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_todo_message_request_proto_rawDescOnce sync.Once
	file_v1_todo_message_request_proto_rawDescData = file_v1_todo_message_request_proto_rawDesc
)

func file_v1_todo_message_request_proto_rawDescGZIP() []byte {
	file_v1_todo_message_request_proto_rawDescOnce.Do(func() {
		file_v1_todo_message_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_todo_message_request_proto_rawDescData)
	})
	return file_v1_todo_message_request_proto_rawDescData
}

var file_v1_todo_message_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_todo_message_request_proto_goTypes = []interface{}{
	(*CreateRequest)(nil), // 0: v1.todo.message.CreateRequest
}
var file_v1_todo_message_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_todo_message_request_proto_init() }
func file_v1_todo_message_request_proto_init() {
	if File_v1_todo_message_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_todo_message_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
			RawDescriptor: file_v1_todo_message_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_todo_message_request_proto_goTypes,
		DependencyIndexes: file_v1_todo_message_request_proto_depIdxs,
		MessageInfos:      file_v1_todo_message_request_proto_msgTypes,
	}.Build()
	File_v1_todo_message_request_proto = out.File
	file_v1_todo_message_request_proto_rawDesc = nil
	file_v1_todo_message_request_proto_goTypes = nil
	file_v1_todo_message_request_proto_depIdxs = nil
}
