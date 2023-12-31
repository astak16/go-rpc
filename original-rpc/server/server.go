package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = rpc.RegisterName("HelloService", &HelloService{})

	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
