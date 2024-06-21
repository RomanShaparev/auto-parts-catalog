// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: auto_part_component_service.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AutoPartComponentServiceClient is the client API for AutoPartComponentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoPartComponentServiceClient interface {
	CreateRootAutoPartComponent(ctx context.Context, in *CreateRootAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error)
	CreateNonRootAutoPartComponent(ctx context.Context, in *CreateNonRootAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error)
	GetAutoPartComponent(ctx context.Context, in *GetAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error)
	UpdateAutoPartComponent(ctx context.Context, in *UpdateAutoPartComponentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAutoPartComponent(ctx context.Context, in *DeleteAutoPartComponentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type autoPartComponentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoPartComponentServiceClient(cc grpc.ClientConnInterface) AutoPartComponentServiceClient {
	return &autoPartComponentServiceClient{cc}
}

func (c *autoPartComponentServiceClient) CreateRootAutoPartComponent(ctx context.Context, in *CreateRootAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error) {
	out := new(AutoPartComponent)
	err := c.cc.Invoke(ctx, "/gen.AutoPartComponentService/CreateRootAutoPartComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoPartComponentServiceClient) CreateNonRootAutoPartComponent(ctx context.Context, in *CreateNonRootAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error) {
	out := new(AutoPartComponent)
	err := c.cc.Invoke(ctx, "/gen.AutoPartComponentService/CreateNonRootAutoPartComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoPartComponentServiceClient) GetAutoPartComponent(ctx context.Context, in *GetAutoPartComponentRequest, opts ...grpc.CallOption) (*AutoPartComponent, error) {
	out := new(AutoPartComponent)
	err := c.cc.Invoke(ctx, "/gen.AutoPartComponentService/GetAutoPartComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoPartComponentServiceClient) UpdateAutoPartComponent(ctx context.Context, in *UpdateAutoPartComponentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gen.AutoPartComponentService/UpdateAutoPartComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoPartComponentServiceClient) DeleteAutoPartComponent(ctx context.Context, in *DeleteAutoPartComponentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/gen.AutoPartComponentService/DeleteAutoPartComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoPartComponentServiceServer is the server API for AutoPartComponentService service.
// All implementations must embed UnimplementedAutoPartComponentServiceServer
// for forward compatibility
type AutoPartComponentServiceServer interface {
	CreateRootAutoPartComponent(context.Context, *CreateRootAutoPartComponentRequest) (*AutoPartComponent, error)
	CreateNonRootAutoPartComponent(context.Context, *CreateNonRootAutoPartComponentRequest) (*AutoPartComponent, error)
	GetAutoPartComponent(context.Context, *GetAutoPartComponentRequest) (*AutoPartComponent, error)
	UpdateAutoPartComponent(context.Context, *UpdateAutoPartComponentRequest) (*emptypb.Empty, error)
	DeleteAutoPartComponent(context.Context, *DeleteAutoPartComponentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAutoPartComponentServiceServer()
}

// UnimplementedAutoPartComponentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAutoPartComponentServiceServer struct {
}

func (UnimplementedAutoPartComponentServiceServer) CreateRootAutoPartComponent(context.Context, *CreateRootAutoPartComponentRequest) (*AutoPartComponent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRootAutoPartComponent not implemented")
}
func (UnimplementedAutoPartComponentServiceServer) CreateNonRootAutoPartComponent(context.Context, *CreateNonRootAutoPartComponentRequest) (*AutoPartComponent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNonRootAutoPartComponent not implemented")
}
func (UnimplementedAutoPartComponentServiceServer) GetAutoPartComponent(context.Context, *GetAutoPartComponentRequest) (*AutoPartComponent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoPartComponent not implemented")
}
func (UnimplementedAutoPartComponentServiceServer) UpdateAutoPartComponent(context.Context, *UpdateAutoPartComponentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAutoPartComponent not implemented")
}
func (UnimplementedAutoPartComponentServiceServer) DeleteAutoPartComponent(context.Context, *DeleteAutoPartComponentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAutoPartComponent not implemented")
}
func (UnimplementedAutoPartComponentServiceServer) mustEmbedUnimplementedAutoPartComponentServiceServer() {
}

// UnsafeAutoPartComponentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoPartComponentServiceServer will
// result in compilation errors.
type UnsafeAutoPartComponentServiceServer interface {
	mustEmbedUnimplementedAutoPartComponentServiceServer()
}

func RegisterAutoPartComponentServiceServer(s grpc.ServiceRegistrar, srv AutoPartComponentServiceServer) {
	s.RegisterService(&AutoPartComponentService_ServiceDesc, srv)
}

func _AutoPartComponentService_CreateRootAutoPartComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRootAutoPartComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoPartComponentServiceServer).CreateRootAutoPartComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.AutoPartComponentService/CreateRootAutoPartComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoPartComponentServiceServer).CreateRootAutoPartComponent(ctx, req.(*CreateRootAutoPartComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoPartComponentService_CreateNonRootAutoPartComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNonRootAutoPartComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoPartComponentServiceServer).CreateNonRootAutoPartComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.AutoPartComponentService/CreateNonRootAutoPartComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoPartComponentServiceServer).CreateNonRootAutoPartComponent(ctx, req.(*CreateNonRootAutoPartComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoPartComponentService_GetAutoPartComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutoPartComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoPartComponentServiceServer).GetAutoPartComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.AutoPartComponentService/GetAutoPartComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoPartComponentServiceServer).GetAutoPartComponent(ctx, req.(*GetAutoPartComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoPartComponentService_UpdateAutoPartComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAutoPartComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoPartComponentServiceServer).UpdateAutoPartComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.AutoPartComponentService/UpdateAutoPartComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoPartComponentServiceServer).UpdateAutoPartComponent(ctx, req.(*UpdateAutoPartComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoPartComponentService_DeleteAutoPartComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAutoPartComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoPartComponentServiceServer).DeleteAutoPartComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.AutoPartComponentService/DeleteAutoPartComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoPartComponentServiceServer).DeleteAutoPartComponent(ctx, req.(*DeleteAutoPartComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoPartComponentService_ServiceDesc is the grpc.ServiceDesc for AutoPartComponentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoPartComponentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gen.AutoPartComponentService",
	HandlerType: (*AutoPartComponentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRootAutoPartComponent",
			Handler:    _AutoPartComponentService_CreateRootAutoPartComponent_Handler,
		},
		{
			MethodName: "CreateNonRootAutoPartComponent",
			Handler:    _AutoPartComponentService_CreateNonRootAutoPartComponent_Handler,
		},
		{
			MethodName: "GetAutoPartComponent",
			Handler:    _AutoPartComponentService_GetAutoPartComponent_Handler,
		},
		{
			MethodName: "UpdateAutoPartComponent",
			Handler:    _AutoPartComponentService_UpdateAutoPartComponent_Handler,
		},
		{
			MethodName: "DeleteAutoPartComponent",
			Handler:    _AutoPartComponentService_DeleteAutoPartComponent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auto_part_component_service.proto",
}
