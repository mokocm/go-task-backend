// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/task/v1/service.proto

package todov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskService_GetTask_FullMethodName    = "/todo.v1.TaskService/GetTask"
	TaskService_GetTasks_FullMethodName   = "/todo.v1.TaskService/GetTasks"
	TaskService_CreateTask_FullMethodName = "/todo.v1.TaskService/CreateTask"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// タスクIDからタスクを取得する
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	// タスク一覧を取得する
	GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksResponse, error)
	// タスクを作成する
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksResponse, error) {
	out := new(GetTasksResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// タスクIDからタスクを取得する
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	// タスク一覧を取得する
	GetTasks(context.Context, *GetTasksRequest) (*GetTasksResponse, error)
	// タスクを作成する
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTasks(context.Context, *GetTasksRequest) (*GetTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedTaskServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTasks(ctx, req.(*GetTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "GetTasks",
			Handler:    _TaskService_GetTasks_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/task/v1/service.proto",
}