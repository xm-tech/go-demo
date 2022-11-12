package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// a rpc method
func (h *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("hello %v", req)
	return nil
}

func Serve() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatal("listen error", err)
	}
	fmt.Println("Server Running")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listen error", err)
		}
		// default use the go blob codec
		go rpc.ServeConn(conn)
	}
}
