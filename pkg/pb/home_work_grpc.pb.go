// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// HomeWorkManagerClient is the client API for HomeWorkManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeWorkManagerClient interface {
	CreateHomeWork(ctx context.Context, in *CreateHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error)
	GetHomeWork(ctx context.Context, in *GetHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error)
	UpdateHomeWork(ctx context.Context, in *UpdateHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error)
	DeleteHomeWork(ctx context.Context, in *DeleteHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error)
	GetHomeWorks(ctx context.Context, in *GetHomeWorksRequest, opts ...grpc.CallOption) (*GetHomeWorksReply, error)
}

type homeWorkManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeWorkManagerClient(cc grpc.ClientConnInterface) HomeWorkManagerClient {
	return &homeWorkManagerClient{cc}
}

func (c *homeWorkManagerClient) CreateHomeWork(ctx context.Context, in *CreateHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error) {
	out := new(HomeWork)
	err := c.cc.Invoke(ctx, "/pb.HomeWorkManager/CreateHomeWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeWorkManagerClient) GetHomeWork(ctx context.Context, in *GetHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error) {
	out := new(HomeWork)
	err := c.cc.Invoke(ctx, "/pb.HomeWorkManager/GetHomeWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeWorkManagerClient) UpdateHomeWork(ctx context.Context, in *UpdateHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error) {
	out := new(HomeWork)
	err := c.cc.Invoke(ctx, "/pb.HomeWorkManager/UpdateHomeWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeWorkManagerClient) DeleteHomeWork(ctx context.Context, in *DeleteHomeWorkRequest, opts ...grpc.CallOption) (*HomeWork, error) {
	out := new(HomeWork)
	err := c.cc.Invoke(ctx, "/pb.HomeWorkManager/DeleteHomeWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeWorkManagerClient) GetHomeWorks(ctx context.Context, in *GetHomeWorksRequest, opts ...grpc.CallOption) (*GetHomeWorksReply, error) {
	out := new(GetHomeWorksReply)
	err := c.cc.Invoke(ctx, "/pb.HomeWorkManager/GetHomeWorks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeWorkManagerServer is the server API for HomeWorkManager service.
// All implementations must embed UnimplementedHomeWorkManagerServer
// for forward compatibility
type HomeWorkManagerServer interface {
	CreateHomeWork(context.Context, *CreateHomeWorkRequest) (*HomeWork, error)
	GetHomeWork(context.Context, *GetHomeWorkRequest) (*HomeWork, error)
	UpdateHomeWork(context.Context, *UpdateHomeWorkRequest) (*HomeWork, error)
	DeleteHomeWork(context.Context, *DeleteHomeWorkRequest) (*HomeWork, error)
	GetHomeWorks(context.Context, *GetHomeWorksRequest) (*GetHomeWorksReply, error)
	mustEmbedUnimplementedHomeWorkManagerServer()
}

// UnimplementedHomeWorkManagerServer must be embedded to have forward compatible implementations.
type UnimplementedHomeWorkManagerServer struct {
}

func (UnimplementedHomeWorkManagerServer) CreateHomeWork(context.Context, *CreateHomeWorkRequest) (*HomeWork, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomeWork not implemented")
}
func (UnimplementedHomeWorkManagerServer) GetHomeWork(context.Context, *GetHomeWorkRequest) (*HomeWork, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeWork not implemented")
}
func (UnimplementedHomeWorkManagerServer) UpdateHomeWork(context.Context, *UpdateHomeWorkRequest) (*HomeWork, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomeWork not implemented")
}
func (UnimplementedHomeWorkManagerServer) DeleteHomeWork(context.Context, *DeleteHomeWorkRequest) (*HomeWork, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHomeWork not implemented")
}
func (UnimplementedHomeWorkManagerServer) GetHomeWorks(context.Context, *GetHomeWorksRequest) (*GetHomeWorksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeWorks not implemented")
}
func (UnimplementedHomeWorkManagerServer) mustEmbedUnimplementedHomeWorkManagerServer() {}

// UnsafeHomeWorkManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeWorkManagerServer will
// result in compilation errors.
type UnsafeHomeWorkManagerServer interface {
	mustEmbedUnimplementedHomeWorkManagerServer()
}

func RegisterHomeWorkManagerServer(s grpc.ServiceRegistrar, srv HomeWorkManagerServer) {
	s.RegisterService(&HomeWorkManager_ServiceDesc, srv)
}

func _HomeWorkManager_CreateHomeWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomeWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeWorkManagerServer).CreateHomeWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HomeWorkManager/CreateHomeWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeWorkManagerServer).CreateHomeWork(ctx, req.(*CreateHomeWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeWorkManager_GetHomeWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeWorkManagerServer).GetHomeWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HomeWorkManager/GetHomeWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeWorkManagerServer).GetHomeWork(ctx, req.(*GetHomeWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeWorkManager_UpdateHomeWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomeWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeWorkManagerServer).UpdateHomeWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HomeWorkManager/UpdateHomeWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeWorkManagerServer).UpdateHomeWork(ctx, req.(*UpdateHomeWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeWorkManager_DeleteHomeWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomeWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeWorkManagerServer).DeleteHomeWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HomeWorkManager/DeleteHomeWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeWorkManagerServer).DeleteHomeWork(ctx, req.(*DeleteHomeWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeWorkManager_GetHomeWorks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeWorksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeWorkManagerServer).GetHomeWorks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HomeWorkManager/GetHomeWorks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeWorkManagerServer).GetHomeWorks(ctx, req.(*GetHomeWorksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeWorkManager_ServiceDesc is the grpc.ServiceDesc for HomeWorkManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeWorkManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HomeWorkManager",
	HandlerType: (*HomeWorkManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHomeWork",
			Handler:    _HomeWorkManager_CreateHomeWork_Handler,
		},
		{
			MethodName: "GetHomeWork",
			Handler:    _HomeWorkManager_GetHomeWork_Handler,
		},
		{
			MethodName: "UpdateHomeWork",
			Handler:    _HomeWorkManager_UpdateHomeWork_Handler,
		},
		{
			MethodName: "DeleteHomeWork",
			Handler:    _HomeWorkManager_DeleteHomeWork_Handler,
		},
		{
			MethodName: "GetHomeWorks",
			Handler:    _HomeWorkManager_GetHomeWorks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home_work.proto",
}
