package main

import (
	"fmt"
	"go-rpc/stream_grpc/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct{}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if data, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(data)
		}
	}
	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到消息：" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{Data: "我是服务器发来的数据"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
