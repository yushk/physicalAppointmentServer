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

// SportItemManagerClient is the client API for SportItemManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SportItemManagerClient interface {
	CreateSportItem(ctx context.Context, in *CreateSportItemRequest, opts ...grpc.CallOption) (*SportItem, error)
	GetSportItem(ctx context.Context, in *GetSportItemRequest, opts ...grpc.CallOption) (*SportItem, error)
	UpdateSportItem(ctx context.Context, in *UpdateSportItemRequest, opts ...grpc.CallOption) (*SportItem, error)
	DeleteSportItem(ctx context.Context, in *DeleteSportItemRequest, opts ...grpc.CallOption) (*SportItem, error)
	GetSportItems(ctx context.Context, in *GetSportItemsRequest, opts ...grpc.CallOption) (*GetSportItemsReply, error)
}

type sportItemManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSportItemManagerClient(cc grpc.ClientConnInterface) SportItemManagerClient {
	return &sportItemManagerClient{cc}
}

func (c *sportItemManagerClient) CreateSportItem(ctx context.Context, in *CreateSportItemRequest, opts ...grpc.CallOption) (*SportItem, error) {
	out := new(SportItem)
	err := c.cc.Invoke(ctx, "/pb.SportItemManager/CreateSportItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportItemManagerClient) GetSportItem(ctx context.Context, in *GetSportItemRequest, opts ...grpc.CallOption) (*SportItem, error) {
	out := new(SportItem)
	err := c.cc.Invoke(ctx, "/pb.SportItemManager/GetSportItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportItemManagerClient) UpdateSportItem(ctx context.Context, in *UpdateSportItemRequest, opts ...grpc.CallOption) (*SportItem, error) {
	out := new(SportItem)
	err := c.cc.Invoke(ctx, "/pb.SportItemManager/UpdateSportItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportItemManagerClient) DeleteSportItem(ctx context.Context, in *DeleteSportItemRequest, opts ...grpc.CallOption) (*SportItem, error) {
	out := new(SportItem)
	err := c.cc.Invoke(ctx, "/pb.SportItemManager/DeleteSportItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportItemManagerClient) GetSportItems(ctx context.Context, in *GetSportItemsRequest, opts ...grpc.CallOption) (*GetSportItemsReply, error) {
	out := new(GetSportItemsReply)
	err := c.cc.Invoke(ctx, "/pb.SportItemManager/GetSportItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SportItemManagerServer is the server API for SportItemManager service.
// All implementations must embed UnimplementedSportItemManagerServer
// for forward compatibility
type SportItemManagerServer interface {
	CreateSportItem(context.Context, *CreateSportItemRequest) (*SportItem, error)
	GetSportItem(context.Context, *GetSportItemRequest) (*SportItem, error)
	UpdateSportItem(context.Context, *UpdateSportItemRequest) (*SportItem, error)
	DeleteSportItem(context.Context, *DeleteSportItemRequest) (*SportItem, error)
	GetSportItems(context.Context, *GetSportItemsRequest) (*GetSportItemsReply, error)
	mustEmbedUnimplementedSportItemManagerServer()
}

// UnimplementedSportItemManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSportItemManagerServer struct {
}

func (UnimplementedSportItemManagerServer) CreateSportItem(context.Context, *CreateSportItemRequest) (*SportItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSportItem not implemented")
}
func (UnimplementedSportItemManagerServer) GetSportItem(context.Context, *GetSportItemRequest) (*SportItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSportItem not implemented")
}
func (UnimplementedSportItemManagerServer) UpdateSportItem(context.Context, *UpdateSportItemRequest) (*SportItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSportItem not implemented")
}
func (UnimplementedSportItemManagerServer) DeleteSportItem(context.Context, *DeleteSportItemRequest) (*SportItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSportItem not implemented")
}
func (UnimplementedSportItemManagerServer) GetSportItems(context.Context, *GetSportItemsRequest) (*GetSportItemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSportItems not implemented")
}
func (UnimplementedSportItemManagerServer) mustEmbedUnimplementedSportItemManagerServer() {}

// UnsafeSportItemManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SportItemManagerServer will
// result in compilation errors.
type UnsafeSportItemManagerServer interface {
	mustEmbedUnimplementedSportItemManagerServer()
}

func RegisterSportItemManagerServer(s grpc.ServiceRegistrar, srv SportItemManagerServer) {
	s.RegisterService(&SportItemManager_ServiceDesc, srv)
}

func _SportItemManager_CreateSportItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSportItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportItemManagerServer).CreateSportItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SportItemManager/CreateSportItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportItemManagerServer).CreateSportItem(ctx, req.(*CreateSportItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SportItemManager_GetSportItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSportItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportItemManagerServer).GetSportItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SportItemManager/GetSportItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportItemManagerServer).GetSportItem(ctx, req.(*GetSportItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SportItemManager_UpdateSportItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSportItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportItemManagerServer).UpdateSportItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SportItemManager/UpdateSportItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportItemManagerServer).UpdateSportItem(ctx, req.(*UpdateSportItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SportItemManager_DeleteSportItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSportItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportItemManagerServer).DeleteSportItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SportItemManager/DeleteSportItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportItemManagerServer).DeleteSportItem(ctx, req.(*DeleteSportItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SportItemManager_GetSportItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSportItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportItemManagerServer).GetSportItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SportItemManager/GetSportItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportItemManagerServer).GetSportItems(ctx, req.(*GetSportItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SportItemManager_ServiceDesc is the grpc.ServiceDesc for SportItemManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SportItemManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SportItemManager",
	HandlerType: (*SportItemManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSportItem",
			Handler:    _SportItemManager_CreateSportItem_Handler,
		},
		{
			MethodName: "GetSportItem",
			Handler:    _SportItemManager_GetSportItem_Handler,
		},
		{
			MethodName: "UpdateSportItem",
			Handler:    _SportItemManager_UpdateSportItem_Handler,
		},
		{
			MethodName: "DeleteSportItem",
			Handler:    _SportItemManager_DeleteSportItem_Handler,
		},
		{
			MethodName: "GetSportItems",
			Handler:    _SportItemManager_GetSportItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template_config.proto",
}
