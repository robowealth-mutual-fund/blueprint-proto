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
	0x74, 0x6f, 0x32, 0x97, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x92,
	0x41, 0x0d, 0x12, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x92, 0x41, 0x14, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0c,
	0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x9e, 0x01, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x6f,
	0x77, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2d, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x2d, 0x66, 0x75,
	0x6e, 0x64, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x92, 0x41,
	0x5c, 0x12, 0x25, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x22, 0x0b, 0x7b, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x50, 0x41, 0x54, 0x48, 0x7d, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v1_todo_service_proto_goTypes = []interface{}{
	(*message.CreateRequest)(nil),  // 0: v1.todo.message.CreateRequest
	(*message.CreateResponse)(nil), // 1: v1.todo.message.CreateResponse
}
var file_v1_todo_service_proto_depIdxs = []int32{
	0, // 0: v1.todo.service.TodoService.Create:input_type -> v1.todo.message.CreateRequest
	1, // 1: v1.todo.service.TodoService.Create:output_type -> v1.todo.message.CreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Create(context.Context, *message.CreateRequest) (*message.CreateResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) Create(context.Context, *message.CreateRequest) (*message.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
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

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.todo.service.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/todo/service.proto",
}
