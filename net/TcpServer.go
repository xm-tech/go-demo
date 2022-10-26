package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	listener, err := net.Listen("tcp4", "127.0.0.1:9999")
	if err != nil {
		log.Panic("net.Listen Fail")
	}
	defer func() {
		listener.Close()
	}()

	go signalHandler()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic("listener.Accept Fail")
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Read Data From Conn Fail, err=%+v", err)
		return
	}
	// echo to client
	conn.Write(buf)
	fmt.Printf("Recevied Data = %+v\n", string(buf[:n]))
}

func signalHandler() {
	c := make(chan os.Signal, 1)
	// 监听系统中断事件，且当中断事件发生时，中断信号写到 os.Signal 类型的 chan 里面
	signal.Notify(c, os.Interrupt)

	for {
		// os.Interrupt event happend
		<-c
		os.Exit(0)
	}
}
