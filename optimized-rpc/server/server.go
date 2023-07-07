package main

import (
	"go-rpc/optimized-rpc/server_proxy"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = server_proxy.RegisterHelloService(&HelloService{})

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
