package main

import (
	"log"
	"net"
	
	"github.com/gin-gonic/examples/grpc/example1/gen/helloworld/v1"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 用于实现 helloworld.GreeterServer。
type server struct {
	v1.UnimplementedGreeterServer
}

// SayHello 实现了 helloworld.GreeterServer 接口

// ff:
// *v1.HelloReply:
// in:
// ctx:
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterGreeterServer(s, &server{})

	// 在gRPC服务器上注册反射服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
