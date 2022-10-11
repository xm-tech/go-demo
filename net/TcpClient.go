// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	input, err := getUserInput(conn)
	fmt.Printf("input = %+v\n", input)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(input))

	go handleResp(conn)

	for {
		//
	}
}

func getUserInput(conn net.Conn) (string, error) {
	fmt.Println("getUserInput")
	scanner := bufio.NewScanner(conn)
	// 循环读取标准输入的数据
	for {
		if !scanner.Scan() {
			err := scanner.Err()
			if err == nil {
				err = io.EOF
			}
			return "", err
		}

		input := scanner.Text()
		return input, nil
	}
}

func handleResp(conn net.Conn) {
	for {
		payload := make([]byte, 512)
		_, err := conn.Read(payload)
		if err != nil {
			conn.Close()
			panic("Read Data Fail, close conn")
		}

		fmt.Printf("resp = %+v\n", string(payload))
	}
}
