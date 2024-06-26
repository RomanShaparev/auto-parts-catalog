// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: warehouse_position_service.proto

package gen

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

// WarehousePositionServiceClient is the client API for WarehousePositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehousePositionServiceClient interface {
	CreateOrUpdateWarehousePosition(ctx context.Context, in *CreateOrUpdateWarehousePositionRequest, opts ...grpc.CallOption) (*WarehousePosition, error)
	GetWarehousePosition(ctx context.Context, in *GetWarehousePositionRequest, opts ...grpc.CallOption) (*WarehousePosition, error)
}

type warehousePositionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehousePositionServiceClient(cc grpc.ClientConnInterface) WarehousePositionServiceClient {
	return &warehousePositionServiceClient{cc}
}

func (c *warehousePositionServiceClient) CreateOrUpdateWarehousePosition(ctx context.Context, in *CreateOrUpdateWarehousePositionRequest, opts ...grpc.CallOption) (*WarehousePosition, error) {
	out := new(WarehousePosition)
	err := c.cc.Invoke(ctx, "/gen.WarehousePositionService/CreateOrUpdateWarehousePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehousePositionServiceClient) GetWarehousePosition(ctx context.Context, in *GetWarehousePositionRequest, opts ...grpc.CallOption) (*WarehousePosition, error) {
	out := new(WarehousePosition)
	err := c.cc.Invoke(ctx, "/gen.WarehousePositionService/GetWarehousePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehousePositionServiceServer is the server API for WarehousePositionService service.
// All implementations must embed UnimplementedWarehousePositionServiceServer
// for forward compatibility
type WarehousePositionServiceServer interface {
	CreateOrUpdateWarehousePosition(context.Context, *CreateOrUpdateWarehousePositionRequest) (*WarehousePosition, error)
	GetWarehousePosition(context.Context, *GetWarehousePositionRequest) (*WarehousePosition, error)
	mustEmbedUnimplementedWarehousePositionServiceServer()
}

// UnimplementedWarehousePositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarehousePositionServiceServer struct {
}

func (UnimplementedWarehousePositionServiceServer) CreateOrUpdateWarehousePosition(context.Context, *CreateOrUpdateWarehousePositionRequest) (*WarehousePosition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateWarehousePosition not implemented")
}
func (UnimplementedWarehousePositionServiceServer) GetWarehousePosition(context.Context, *GetWarehousePositionRequest) (*WarehousePosition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehousePosition not implemented")
}
func (UnimplementedWarehousePositionServiceServer) mustEmbedUnimplementedWarehousePositionServiceServer() {
}

// UnsafeWarehousePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehousePositionServiceServer will
// result in compilation errors.
type UnsafeWarehousePositionServiceServer interface {
	mustEmbedUnimplementedWarehousePositionServiceServer()
}

func RegisterWarehousePositionServiceServer(s grpc.ServiceRegistrar, srv WarehousePositionServiceServer) {
	s.RegisterService(&WarehousePositionService_ServiceDesc, srv)
}

func _WarehousePositionService_CreateOrUpdateWarehousePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateWarehousePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehousePositionServiceServer).CreateOrUpdateWarehousePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.WarehousePositionService/CreateOrUpdateWarehousePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehousePositionServiceServer).CreateOrUpdateWarehousePosition(ctx, req.(*CreateOrUpdateWarehousePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehousePositionService_GetWarehousePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarehousePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehousePositionServiceServer).GetWarehousePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.WarehousePositionService/GetWarehousePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehousePositionServiceServer).GetWarehousePosition(ctx, req.(*GetWarehousePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WarehousePositionService_ServiceDesc is the grpc.ServiceDesc for WarehousePositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarehousePositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gen.WarehousePositionService",
	HandlerType: (*WarehousePositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateWarehousePosition",
			Handler:    _WarehousePositionService_CreateOrUpdateWarehousePosition_Handler,
		},
		{
			MethodName: "GetWarehousePosition",
			Handler:    _WarehousePositionService_GetWarehousePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse_position_service.proto",
}
