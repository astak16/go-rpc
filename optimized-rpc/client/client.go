package main

import (
	"fmt"
	"go-rpc/optimized-rpc/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	var require *string = new(string)
	err := client.Hello("uccs", require)
	if err != nil {
		panic(err)
	}
	fmt.Println(*require)
}
