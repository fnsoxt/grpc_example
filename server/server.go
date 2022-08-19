package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	service "hello_grpc/service"
	"net"
	"net/http"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloReply, error) {
	return &service.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8080")
	fmt.Println("start server :8080")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	service.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	go func() {
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("开启服务失败: %s", err)
			return
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = service.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		fmt.Println("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	fmt.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	fmt.Println(gwServer.ListenAndServe())
}
