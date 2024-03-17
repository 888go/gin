// 由protoc-gen-go-grpc生成的代码
// 不要编辑
// 版本号:- protoc-gen-go-grpc v1.2.0 - protoc v3.21.5

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// 这是一个编译时断言，用于确保生成的文件与正在对其进行编译的grpc包兼容
// 要求gRPC-Go v1.32.0或更高版本
const _ = grpc.SupportPackageIsVersion7

// GreeterClient是greter服务的客户端API
// 关于ctx使用和关闭/结束流rpc的语义，请参阅https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream
type GreeterClient interface {
// 发送问候
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}


// ff:
// cc:

// ff:
// cc:
func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}


// ff:
// *HelloReply:
// opts:
// in:
// ctx:

// ff:
// *HelloReply:
// opts:
// in:
// ctx:
func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer是greter服务的服务器API
// 为了向前兼容，所有实现必须嵌入unimplementgreeterserver
type GreeterServer interface {
// 发送问候
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer必须被嵌入以具有向前兼容的实现
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer可以内嵌，以选择退出此服务的前向兼容性
// 不建议使用此接口，因为向GreeterServer添加的方法将导致编译错误
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}


// ff:
// srv:
// s:

// ff:
// srv:
// s:
func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc是grpc
// ServiceDesc表示迎宾服务
// 它只适合与grpc直接使用
// RegisterService，并且不能被自省或修改(即使作为副本)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/helloworld/v1/helloworld.proto",
}
