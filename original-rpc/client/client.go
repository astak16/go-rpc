package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var require *string = new(string)
	err = conn.Call("HelloService.Hello", "uccs", require)
	if err != nil {
		panic(err)
	}
	fmt.Println(*require)
}
