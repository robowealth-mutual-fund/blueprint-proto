// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.0
// source: v1/todo/service.proto

package todo

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	message "github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_todo_service_proto protoreflect.FileDescriptor

var file_v1_todo_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xd4, 0x04, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x92,
	0x41, 0x0d, 0x12, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x92, 0x41, 0x11, 0x12, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12,
	0x6b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x92, 0x41, 0x10, 0x12, 0x0e, 0x47, 0x65, 0x74, 0x20, 0x74, 0x6f, 0x64, 0x6f,
	0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6f, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x92, 0x41, 0x0d, 0x12, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x1a,
	0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x92, 0x41, 0x0d, 0x12, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x1a, 0x17, 0x92, 0x41, 0x14, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0c, 0x54, 0x6f, 0x64,
	0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x9e, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x77, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2d, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x92, 0x41, 0x5c, 0x12, 0x25,
	0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x05,
	0x31, 0x2e, 0x30, 0x2e, 0x31, 0x22, 0x0b, 0x7b, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x41, 0x54,
	0x48, 0x7d, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_v1_todo_service_proto_goTypes = []interface{}{
	(*message.CreateRequest)(nil),  // 0: v1.todo.message.CreateRequest
	(*message.ListRequest)(nil),    // 1: v1.todo.message.ListRequest
	(*message.GetRequest)(nil),     // 2: v1.todo.message.GetRequest
	(*message.UpdateRequest)(nil),  // 3: v1.todo.message.UpdateRequest
	(*message.DeleteRequest)(nil),  // 4: v1.todo.message.DeleteRequest
	(*message.CreateResponse)(nil), // 5: v1.todo.message.CreateResponse
	(*message.ListResponse)(nil),   // 6: v1.todo.message.ListResponse
	(*message.GetResponse)(nil),    // 7: v1.todo.message.GetResponse
	(*message.UpdateResponse)(nil), // 8: v1.todo.message.UpdateResponse
	(*message.DeleteResponse)(nil), // 9: v1.todo.message.DeleteResponse
}
var file_v1_todo_service_proto_depIdxs = []int32{
	0, // 0: v1.todo.service.TodoService.Create:input_type -> v1.todo.message.CreateRequest
	1, // 1: v1.todo.service.TodoService.List:input_type -> v1.todo.message.ListRequest
	2, // 2: v1.todo.service.TodoService.Get:input_type -> v1.todo.message.GetRequest
	3, // 3: v1.todo.service.TodoService.Update:input_type -> v1.todo.message.UpdateRequest
	4, // 4: v1.todo.service.TodoService.Delete:input_type -> v1.todo.message.DeleteRequest
	5, // 5: v1.todo.service.TodoService.Create:output_type -> v1.todo.message.CreateResponse
	6, // 6: v1.todo.service.TodoService.List:output_type -> v1.todo.message.ListResponse
	7, // 7: v1.todo.service.TodoService.Get:output_type -> v1.todo.message.GetResponse
	8, // 8: v1.todo.service.TodoService.Update:output_type -> v1.todo.message.UpdateResponse
	9, // 9: v1.todo.service.TodoService.Delete:output_type -> v1.todo.message.DeleteResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_todo_service_proto_init() }
func file_v1_todo_service_proto_init() {
	if File_v1_todo_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_todo_service_proto_goTypes,
		DependencyIndexes: file_v1_todo_service_proto_depIdxs,
	}.Build()
	File_v1_todo_service_proto = out.File
	file_v1_todo_service_proto_rawDesc = nil
	file_v1_todo_service_proto_goTypes = nil
	file_v1_todo_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	Create(ctx context.Context, in *message.CreateRequest, opts ...grpc.CallOption) (*message.CreateResponse, error)
	List(ctx context.Context, in *message.ListRequest, opts ...grpc.CallOption) (*message.ListResponse, error)
	Get(ctx context.Context, in *message.GetRequest, opts ...grpc.CallOption) (*message.GetResponse, error)
	Update(ctx context.Context, in *message.UpdateRequest, opts ...grpc.CallOption) (*message.UpdateResponse, error)
	Delete(ctx context.Context, in *message.DeleteRequest, opts ...grpc.CallOption) (*message.DeleteResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Create(ctx context.Context, in *message.CreateRequest, opts ...grpc.CallOption) (*message.CreateResponse, error) {
	out := new(message.CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.todo.service.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) List(ctx context.Context, in *message.ListRequest, opts ...grpc.CallOption) (*message.ListResponse, error) {
	out := new(message.ListResponse)
	err := c.cc.Invoke(ctx, "/v1.todo.service.TodoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Get(ctx context.Context, in *message.GetRequest, opts ...grpc.CallOption) (*message.GetResponse, error) {
	out := new(message.GetResponse)
	err := c.cc.Invoke(ctx, "/v1.todo.service.TodoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *message.UpdateRequest, opts ...grpc.CallOption) (*message.UpdateResponse, error) {
	out := new(message.UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.todo.service.TodoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *message.DeleteRequest, opts ...grpc.CallOption) (*message.DeleteResponse, error) {
	out := new(message.DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.todo.service.TodoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Create(context.Context, *message.CreateRequest) (*message.CreateResponse, error)
	List(context.Context, *message.ListRequest) (*message.ListResponse, error)
	Get(context.Context, *message.GetRequest) (*message.GetResponse, error)
	Update(context.Context, *message.UpdateRequest) (*message.UpdateResponse, error)
	Delete(context.Context, *message.DeleteRequest) (*message.DeleteResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) Create(context.Context, *message.CreateRequest) (*message.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTodoServiceServer) List(context.Context, *message.ListRequest) (*message.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTodoServiceServer) Get(context.Context, *message.GetRequest) (*message.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTodoServiceServer) Update(context.Context, *message.UpdateRequest) (*message.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTodoServiceServer) Delete(context.Context, *message.DeleteRequest) (*message.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.todo.service.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*message.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.todo.service.TodoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).List(ctx, req.(*message.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.todo.service.TodoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Get(ctx, req.(*message.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.todo.service.TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*message.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.todo.service.TodoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*message.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.todo.service.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TodoService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TodoService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/todo/service.proto",
}
