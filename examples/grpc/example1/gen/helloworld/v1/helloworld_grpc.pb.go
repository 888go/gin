// 由protoc-gen-go-grpc工具生成的代码。请勿编辑。
// 版本信息：
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// 源文件：pb/helloworld/v1/helloworld.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// 这是一个编译时断言，用于确保此生成文件
// 与正在编译的gRPC包兼容。
// 需要gRPC-Go v1.32.0或更高版本。
const _ = grpc.SupportPackageIsVersion7

// GreeterClient 是 Greeter 服务的客户端 API。
//
// 关于 ctx 的使用语义以及关闭/结束流式 RPC，请参考 https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream。
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}


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
func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer 是 Greeter 服务的服务器端 API。
// 所有实现都必须嵌入 UnimplementedGreeterServer，以保证向前兼容性
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer 需要被嵌入以确保兼容未来的实现。
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer 可以被嵌入以选择性地退出此服务的向前兼容性。
// 不建议使用此接口，因为向 GreeterServer 添加方法会导致编译错误。
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}


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

// Greeter_ServiceDesc 是 Greeter 服务的 grpc.ServiceDesc。
// 它仅用于直接配合 grpc.RegisterService 使用，
// 不应被深入检查或修改（即使是作为副本）
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
