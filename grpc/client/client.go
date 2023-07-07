package main

import (
	"context"
	"fmt"
	"go-rpc/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	r, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "uccs"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}
