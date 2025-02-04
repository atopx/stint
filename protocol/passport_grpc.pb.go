// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: protocol/passport.proto

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PassportService_Health_FullMethodName = "/protocol.PassportService/Health"
	PassportService_Shell_FullMethodName  = "/protocol.PassportService/Shell"
)

// PassportServiceClient is the client API for PassportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PassportService 认证服务
type PassportServiceClient interface {
	// Health 服务健康检查
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Shell 交互命令, 第一次需要使用token建立握手
	Shell(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error)
}

type passportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPassportServiceClient(cc grpc.ClientConnInterface) PassportServiceClient {
	return &passportServiceClient{cc}
}

func (c *passportServiceClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PassportService_Health_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportServiceClient) Shell(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PassportService_ServiceDesc.Streams[0], PassportService_Shell_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Message, Message]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PassportService_ShellClient = grpc.BidiStreamingClient[Message, Message]

// PassportServiceServer is the server API for PassportService service.
// All implementations must embed UnimplementedPassportServiceServer
// for forward compatibility.
//
// PassportService 认证服务
type PassportServiceServer interface {
	// Health 服务健康检查
	Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Shell 交互命令, 第一次需要使用token建立握手
	Shell(grpc.BidiStreamingServer[Message, Message]) error
	mustEmbedUnimplementedPassportServiceServer()
}

// UnimplementedPassportServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPassportServiceServer struct{}

func (UnimplementedPassportServiceServer) Health(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedPassportServiceServer) Shell(grpc.BidiStreamingServer[Message, Message]) error {
	return status.Errorf(codes.Unimplemented, "method Shell not implemented")
}
func (UnimplementedPassportServiceServer) mustEmbedUnimplementedPassportServiceServer() {}
func (UnimplementedPassportServiceServer) testEmbeddedByValue()                         {}

// UnsafePassportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassportServiceServer will
// result in compilation errors.
type UnsafePassportServiceServer interface {
	mustEmbedUnimplementedPassportServiceServer()
}

func RegisterPassportServiceServer(s grpc.ServiceRegistrar, srv PassportServiceServer) {
	// If the following call pancis, it indicates UnimplementedPassportServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PassportService_ServiceDesc, srv)
}

func _PassportService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PassportService_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServiceServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassportService_Shell_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PassportServiceServer).Shell(&grpc.GenericServerStream[Message, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PassportService_ShellServer = grpc.BidiStreamingServer[Message, Message]

// PassportService_ServiceDesc is the grpc.ServiceDesc for PassportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PassportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.PassportService",
	HandlerType: (*PassportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _PassportService_Health_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Shell",
			Handler:       _PassportService_Shell_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protocol/passport.proto",
}
