// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: star.proto

package api

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

// StarConfigClient is the client API for StarConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StarConfigClient interface {
	GetStandaloneConfig(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*NodeStandaloneConfig, error)
	GetConfigGroup(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*NodeConfigGroup, error)
}

type starConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewStarConfigClient(cc grpc.ClientConnInterface) StarConfigClient {
	return &starConfigClient{cc}
}

func (c *starConfigClient) GetStandaloneConfig(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*NodeStandaloneConfig, error) {
	out := new(NodeStandaloneConfig)
	err := c.cc.Invoke(ctx, "/proto.StarConfig/GetStandaloneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starConfigClient) GetConfigGroup(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*NodeConfigGroup, error) {
	out := new(NodeConfigGroup)
	err := c.cc.Invoke(ctx, "/proto.StarConfig/GetConfigGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarConfigServer is the server API for StarConfig service.
// All implementations must embed UnimplementedStarConfigServer
// for forward compatibility
type StarConfigServer interface {
	GetStandaloneConfig(context.Context, *GetReq) (*NodeStandaloneConfig, error)
	GetConfigGroup(context.Context, *GetReq) (*NodeConfigGroup, error)
	mustEmbedUnimplementedStarConfigServer()
}

// UnimplementedStarConfigServer must be embedded to have forward compatible implementations.
type UnimplementedStarConfigServer struct {
}

func (UnimplementedStarConfigServer) GetStandaloneConfig(context.Context, *GetReq) (*NodeStandaloneConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandaloneConfig not implemented")
}
func (UnimplementedStarConfigServer) GetConfigGroup(context.Context, *GetReq) (*NodeConfigGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigGroup not implemented")
}
func (UnimplementedStarConfigServer) mustEmbedUnimplementedStarConfigServer() {}

// UnsafeStarConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StarConfigServer will
// result in compilation errors.
type UnsafeStarConfigServer interface {
	mustEmbedUnimplementedStarConfigServer()
}

func RegisterStarConfigServer(s grpc.ServiceRegistrar, srv StarConfigServer) {
	s.RegisterService(&StarConfig_ServiceDesc, srv)
}

func _StarConfig_GetStandaloneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarConfigServer).GetStandaloneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StarConfig/GetStandaloneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarConfigServer).GetStandaloneConfig(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarConfig_GetConfigGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarConfigServer).GetConfigGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StarConfig/GetConfigGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarConfigServer).GetConfigGroup(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StarConfig_ServiceDesc is the grpc.ServiceDesc for StarConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StarConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StarConfig",
	HandlerType: (*StarConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStandaloneConfig",
			Handler:    _StarConfig_GetStandaloneConfig_Handler,
		},
		{
			MethodName: "GetConfigGroup",
			Handler:    _StarConfig_GetConfigGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "star.proto",
}
