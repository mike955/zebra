// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package flake

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

// FlakeServiceClient is the client API for FlakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlakeServiceClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error)
}

type flakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlakeServiceClient(cc grpc.ClientConnInterface) FlakeServiceClient {
	return &flakeServiceClient{cc}
}

func (c *flakeServiceClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error) {
	out := new(NewResponse)
	err := c.cc.Invoke(ctx, "/flake.FlakeService/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlakeServiceServer is the server API for FlakeService service.
// All implementations must embed UnimplementedFlakeServiceServer
// for forward compatibility
type FlakeServiceServer interface {
	New(context.Context, *NewRequest) (*NewResponse, error)
	mustEmbedUnimplementedFlakeServiceServer()
}

// UnimplementedFlakeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlakeServiceServer struct {
}

func (UnimplementedFlakeServiceServer) New(context.Context, *NewRequest) (*NewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedFlakeServiceServer) mustEmbedUnimplementedFlakeServiceServer() {}

// UnsafeFlakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlakeServiceServer will
// result in compilation errors.
type UnsafeFlakeServiceServer interface {
	mustEmbedUnimplementedFlakeServiceServer()
}

func RegisterFlakeServiceServer(s grpc.ServiceRegistrar, srv FlakeServiceServer) {
	s.RegisterService(&FlakeService_ServiceDesc, srv)
}

func _FlakeService_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlakeServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flake.FlakeService/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlakeServiceServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlakeService_ServiceDesc is the grpc.ServiceDesc for FlakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flake.FlakeService",
	HandlerType: (*FlakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _FlakeService_New_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flake.proto",
}
