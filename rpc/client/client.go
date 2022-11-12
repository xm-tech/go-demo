package client

import (
	"fmt"
	"log"
	"net/rpc"
)

func Run() {
	client, err := rpc.Dial("tcp", ":1111")
	if err != nil {
		log.Fatal("rpc.Dial error", err)
	}

	fmt.Println("Client call: HelloService.Hello, Param:golang")
	var resp string
	err = client.Call("HelloService.Hello", "golang", &resp)
	if err != nil {
		log.Fatal("client.Call error", err)
	}
	fmt.Println("Server resp:", resp)
}
