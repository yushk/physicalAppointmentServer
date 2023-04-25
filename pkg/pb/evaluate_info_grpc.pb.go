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

// EvaluateInfoManagerClient is the client API for EvaluateInfoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvaluateInfoManagerClient interface {
	CreateEvaluateInfo(ctx context.Context, in *CreateEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error)
	UpdateEvaluateInfo(ctx context.Context, in *UpdateEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error)
	DeleteEvaluateInfo(ctx context.Context, in *DeleteEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error)
	GetEvaluateInfos(ctx context.Context, in *GetEvaluateInfosRequest, opts ...grpc.CallOption) (*GetEvaluateInfosReply, error)
}

type evaluateInfoManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewEvaluateInfoManagerClient(cc grpc.ClientConnInterface) EvaluateInfoManagerClient {
	return &evaluateInfoManagerClient{cc}
}

func (c *evaluateInfoManagerClient) CreateEvaluateInfo(ctx context.Context, in *CreateEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error) {
	out := new(EvaluateInfo)
	err := c.cc.Invoke(ctx, "/pb.EvaluateInfoManager/CreateEvaluateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluateInfoManagerClient) UpdateEvaluateInfo(ctx context.Context, in *UpdateEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error) {
	out := new(EvaluateInfo)
	err := c.cc.Invoke(ctx, "/pb.EvaluateInfoManager/UpdateEvaluateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluateInfoManagerClient) DeleteEvaluateInfo(ctx context.Context, in *DeleteEvaluateInfoRequest, opts ...grpc.CallOption) (*EvaluateInfo, error) {
	out := new(EvaluateInfo)
	err := c.cc.Invoke(ctx, "/pb.EvaluateInfoManager/DeleteEvaluateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluateInfoManagerClient) GetEvaluateInfos(ctx context.Context, in *GetEvaluateInfosRequest, opts ...grpc.CallOption) (*GetEvaluateInfosReply, error) {
	out := new(GetEvaluateInfosReply)
	err := c.cc.Invoke(ctx, "/pb.EvaluateInfoManager/GetEvaluateInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluateInfoManagerServer is the server API for EvaluateInfoManager service.
// All implementations must embed UnimplementedEvaluateInfoManagerServer
// for forward compatibility
type EvaluateInfoManagerServer interface {
	CreateEvaluateInfo(context.Context, *CreateEvaluateInfoRequest) (*EvaluateInfo, error)
	UpdateEvaluateInfo(context.Context, *UpdateEvaluateInfoRequest) (*EvaluateInfo, error)
	DeleteEvaluateInfo(context.Context, *DeleteEvaluateInfoRequest) (*EvaluateInfo, error)
	GetEvaluateInfos(context.Context, *GetEvaluateInfosRequest) (*GetEvaluateInfosReply, error)
	mustEmbedUnimplementedEvaluateInfoManagerServer()
}

// UnimplementedEvaluateInfoManagerServer must be embedded to have forward compatible implementations.
type UnimplementedEvaluateInfoManagerServer struct {
}

func (UnimplementedEvaluateInfoManagerServer) CreateEvaluateInfo(context.Context, *CreateEvaluateInfoRequest) (*EvaluateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvaluateInfo not implemented")
}
func (UnimplementedEvaluateInfoManagerServer) UpdateEvaluateInfo(context.Context, *UpdateEvaluateInfoRequest) (*EvaluateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvaluateInfo not implemented")
}
func (UnimplementedEvaluateInfoManagerServer) DeleteEvaluateInfo(context.Context, *DeleteEvaluateInfoRequest) (*EvaluateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvaluateInfo not implemented")
}
func (UnimplementedEvaluateInfoManagerServer) GetEvaluateInfos(context.Context, *GetEvaluateInfosRequest) (*GetEvaluateInfosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvaluateInfos not implemented")
}
func (UnimplementedEvaluateInfoManagerServer) mustEmbedUnimplementedEvaluateInfoManagerServer() {}

// UnsafeEvaluateInfoManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvaluateInfoManagerServer will
// result in compilation errors.
type UnsafeEvaluateInfoManagerServer interface {
	mustEmbedUnimplementedEvaluateInfoManagerServer()
}

func RegisterEvaluateInfoManagerServer(s grpc.ServiceRegistrar, srv EvaluateInfoManagerServer) {
	s.RegisterService(&EvaluateInfoManager_ServiceDesc, srv)
}

func _EvaluateInfoManager_CreateEvaluateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEvaluateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluateInfoManagerServer).CreateEvaluateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EvaluateInfoManager/CreateEvaluateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluateInfoManagerServer).CreateEvaluateInfo(ctx, req.(*CreateEvaluateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluateInfoManager_UpdateEvaluateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEvaluateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluateInfoManagerServer).UpdateEvaluateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EvaluateInfoManager/UpdateEvaluateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluateInfoManagerServer).UpdateEvaluateInfo(ctx, req.(*UpdateEvaluateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluateInfoManager_DeleteEvaluateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEvaluateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluateInfoManagerServer).DeleteEvaluateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EvaluateInfoManager/DeleteEvaluateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluateInfoManagerServer).DeleteEvaluateInfo(ctx, req.(*DeleteEvaluateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluateInfoManager_GetEvaluateInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEvaluateInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluateInfoManagerServer).GetEvaluateInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EvaluateInfoManager/GetEvaluateInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluateInfoManagerServer).GetEvaluateInfos(ctx, req.(*GetEvaluateInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EvaluateInfoManager_ServiceDesc is the grpc.ServiceDesc for EvaluateInfoManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EvaluateInfoManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EvaluateInfoManager",
	HandlerType: (*EvaluateInfoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvaluateInfo",
			Handler:    _EvaluateInfoManager_CreateEvaluateInfo_Handler,
		},
		{
			MethodName: "UpdateEvaluateInfo",
			Handler:    _EvaluateInfoManager_UpdateEvaluateInfo_Handler,
		},
		{
			MethodName: "DeleteEvaluateInfo",
			Handler:    _EvaluateInfoManager_DeleteEvaluateInfo_Handler,
		},
		{
			MethodName: "GetEvaluateInfos",
			Handler:    _EvaluateInfoManager_GetEvaluateInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evaluate_info.proto",
}