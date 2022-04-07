package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	service "hello_grpc/service"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("连接服务端失败：%s", err)
		return
	}

	defer conn.Close()

	c := service.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &service.HelloRequest{Name: "horika"})
	if err != nil {
		fmt.Printf("调用服务端代码失败：%s", err)
		return
	}

	fmt.Printf("调用成功：%s", r.Message)
}
