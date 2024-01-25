// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: log_compression.proto

package logcompressionpb

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
	LogCompressionService_CompressLog_FullMethodName = "/logcompression.LogCompressionService/CompressLog"
)

// LogCompressionServiceClient is the client API for LogCompressionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogCompressionServiceClient interface {
	CompressLog(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*CompressionMapping, error)
}

type logCompressionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogCompressionServiceClient(cc grpc.ClientConnInterface) LogCompressionServiceClient {
	return &logCompressionServiceClient{cc}
}

func (c *logCompressionServiceClient) CompressLog(ctx context.Context, in *LogMessage, opts ...grpc.CallOption) (*CompressionMapping, error) {
	out := new(CompressionMapping)
	err := c.cc.Invoke(ctx, LogCompressionService_CompressLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogCompressionServiceServer is the server API for LogCompressionService service.
// All implementations must embed UnimplementedLogCompressionServiceServer
// for forward compatibility
type LogCompressionServiceServer interface {
	CompressLog(context.Context, *LogMessage) (*CompressionMapping, error)
	mustEmbedUnimplementedLogCompressionServiceServer()
}

// UnimplementedLogCompressionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogCompressionServiceServer struct {
}

func (UnimplementedLogCompressionServiceServer) CompressLog(context.Context, *LogMessage) (*CompressionMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompressLog not implemented")
}
func (UnimplementedLogCompressionServiceServer) mustEmbedUnimplementedLogCompressionServiceServer() {}

// UnsafeLogCompressionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogCompressionServiceServer will
// result in compilation errors.
type UnsafeLogCompressionServiceServer interface {
	mustEmbedUnimplementedLogCompressionServiceServer()
}

func RegisterLogCompressionServiceServer(s grpc.ServiceRegistrar, srv LogCompressionServiceServer) {
	s.RegisterService(&LogCompressionService_ServiceDesc, srv)
}

func _LogCompressionService_CompressLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogCompressionServiceServer).CompressLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogCompressionService_CompressLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogCompressionServiceServer).CompressLog(ctx, req.(*LogMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// LogCompressionService_ServiceDesc is the grpc.ServiceDesc for LogCompressionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogCompressionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logcompression.LogCompressionService",
	HandlerType: (*LogCompressionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompressLog",
			Handler:    _LogCompressionService_CompressLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log_compression.proto",
}
