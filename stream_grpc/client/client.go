package main

import (
	"context"
	"fmt"
	"go-rpc/stream_grpc/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func serverSteam(c proto.GreeterClient) {
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "uccs",
	})
	for {
		data, err := res.Recv() // 服务端用 send 发送，客户端用 recv 接收
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(data)
	}
}
func clientSteam(c proto.GreeterClient) {
	putRes, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putRes.Send(&proto.StreamReqData{
			Data: "uccs",
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}

func allSteam(c proto.GreeterClient) {
	allStr, _ := c.AllStream(context.Background())
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
			_ = allStr.Send(&proto.StreamReqData{Data: "我是uccs"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	serverSteam(c)
	clientSteam(c)
	allSteam(c)

}
