package client_proxy

import (
	"go-rpc/optimized-rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol string, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("连接失败")
	}
	return HelloServiceStub{
		conn,
	}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	return c.Client.Call(handler.HelloServiceName+".Hello", request, reply)
}
