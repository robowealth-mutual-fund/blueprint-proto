// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.0
// source: v1/todo/message/response.proto

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

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName  string `protobuf:"bytes,1,opt,name=taskName,proto3" json:"taskName,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_response_proto_rawDescGZIP(), []int{0}
}

func (x *CreateResponse) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *CreateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CreateResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	TaskName  []string `protobuf:"bytes,2,rep,name=taskName,proto3" json:"taskName,omitempty"`
	Status    []string `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
	CreatedAt []int64  `protobuf:"varint,4,rep,packed,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt []int64  `protobuf:"varint,5,rep,packed,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_response_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListResponse) GetTaskName() []string {
	if x != nil {
		return x.TaskName
	}
	return nil
}

func (x *ListResponse) GetStatus() []string {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListResponse) GetCreatedAt() []int64 {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListResponse) GetUpdatedAt() []int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskName  string `protobuf:"bytes,2,opt,name=taskName,proto3" json:"taskName,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_response_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetResponse) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *GetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_response_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optionally include a confirmation message or other relevant info
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_todo_message_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todo_message_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_v1_todo_message_response_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_todo_message_response_proto protoreflect.FileDescriptor

var file_v1_todo_message_response_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2,
	0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x4a, 0x0c, 0x22, 0x53, 0x65, 0x6e, 0x64, 0x20,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x54, 0x4f, 0x44, 0x4f, 0x22, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31,
	0x37, 0x31, 0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31, 0x37,
	0x31, 0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26, 0x22, 0x33, 0x36, 0x61, 0x61, 0x34, 0x31, 0x36, 0x36,
	0x2d, 0x38, 0x61, 0x32, 0x37, 0x2d, 0x34, 0x32, 0x35, 0x66, 0x2d, 0x62, 0x39, 0x61, 0x37, 0x2d,
	0x34, 0x63, 0x64, 0x31, 0x62, 0x34, 0x64, 0x36, 0x35, 0x36, 0x32, 0x39, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x4a, 0x0c, 0x22, 0x53, 0x65, 0x6e, 0x64, 0x20,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x54, 0x4f, 0x44, 0x4f, 0x22, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31,
	0x37, 0x31, 0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31, 0x37,
	0x31, 0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26, 0x22, 0x33, 0x36, 0x61, 0x61, 0x34, 0x31, 0x36, 0x36, 0x2d,
	0x38, 0x61, 0x32, 0x37, 0x2d, 0x34, 0x32, 0x35, 0x66, 0x2d, 0x62, 0x39, 0x61, 0x37, 0x2d, 0x34,
	0x63, 0x64, 0x31, 0x62, 0x34, 0x64, 0x36, 0x35, 0x36, 0x32, 0x39, 0x22, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x4a, 0x0c, 0x22, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x54, 0x4f, 0x44, 0x4f, 0x22, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31, 0x37,
	0x31, 0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x31, 0x37, 0x31,
	0x34, 0x33, 0x36, 0x39, 0x39, 0x30, 0x34, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26, 0x22, 0x33, 0x36, 0x61, 0x61, 0x34, 0x31, 0x36, 0x36,
	0x2d, 0x38, 0x61, 0x32, 0x37, 0x2d, 0x34, 0x32, 0x35, 0x66, 0x2d, 0x62, 0x39, 0x61, 0x37, 0x2d,
	0x34, 0x63, 0x64, 0x31, 0x62, 0x34, 0x64, 0x36, 0x35, 0x36, 0x32, 0x39, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x51, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x92, 0x41, 0x22, 0x4a, 0x20, 0x22, 0x54, 0x6f, 0x64, 0x6f,
	0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x77, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2d, 0x6d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_todo_message_response_proto_rawDescOnce sync.Once
	file_v1_todo_message_response_proto_rawDescData = file_v1_todo_message_response_proto_rawDesc
)

func file_v1_todo_message_response_proto_rawDescGZIP() []byte {
	file_v1_todo_message_response_proto_rawDescOnce.Do(func() {
		file_v1_todo_message_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_todo_message_response_proto_rawDescData)
	})
	return file_v1_todo_message_response_proto_rawDescData
}

var file_v1_todo_message_response_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_todo_message_response_proto_goTypes = []interface{}{
	(*CreateResponse)(nil), // 0: v1.todo.message.CreateResponse
	(*ListResponse)(nil),   // 1: v1.todo.message.ListResponse
	(*GetResponse)(nil),    // 2: v1.todo.message.GetResponse
	(*UpdateResponse)(nil), // 3: v1.todo.message.UpdateResponse
	(*DeleteResponse)(nil), // 4: v1.todo.message.DeleteResponse
}
var file_v1_todo_message_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_todo_message_response_proto_init() }
func file_v1_todo_message_response_proto_init() {
	if File_v1_todo_message_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_todo_message_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_v1_todo_message_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_v1_todo_message_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_v1_todo_message_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_v1_todo_message_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_v1_todo_message_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_todo_message_response_proto_goTypes,
		DependencyIndexes: file_v1_todo_message_response_proto_depIdxs,
		MessageInfos:      file_v1_todo_message_response_proto_msgTypes,
	}.Build()
	File_v1_todo_message_response_proto = out.File
	file_v1_todo_message_response_proto_rawDesc = nil
	file_v1_todo_message_response_proto_goTypes = nil
	file_v1_todo_message_response_proto_depIdxs = nil
}
