package main

import (
	"context"
	"go-rpc/grpc/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := g.Serve(lis); err != nil {
		panic(err)
	}
}
