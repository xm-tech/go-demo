package main

import (
	"time"

	"github.com/xm-tech/go-demo/rpc/client"
	"github.com/xm-tech/go-demo/rpc/server"
)

func main() {
	go server.Serve()

	time.Sleep(3 * time.Second)

	go client.Run()

	select {}
}
