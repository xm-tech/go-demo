package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

func main() {
	delete_file("/Volumes/My Passport/Downloads")

	// other_op()
}

func delete_file(path string) {
	fmt.Println("try to delete_file:", path)
	if err := os.RemoveAll(path); err != nil {
		log.Println("delete_file fail, path=", path)
		panic(err)
	}
}

func other_op() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic("create file fail")
	}

	defer file.Close()
	fmt.Printf("file = %+v\n", file)

	num, err := file.WriteString("hello, golang")
	if err != nil {
		panic("write file fail")
	}
	fmt.Printf("num = %+v\n", num)

	file, err = os.Open("text.txt")
	if err != nil {
		panic("open file fail")
	}
	fmt.Printf("file = %+v\n", file)

	// read 2 bytes, read pointer moved
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		panic("read fail")
	}
	fmt.Printf("numBytesRead = %+v, byteSlice = %+v\n", numBytesRead, byteSlice)

	// read all
	bytesRead, err := ioutil.ReadAll(file)
	if err != nil {
		panic("ioutil.ReadAll fail")
	}
	fmt.Printf("bytesRead = %+v\n", bytesRead)

	// bytes to string
	str := string(bytesRead[:])
	fmt.Printf("str = %+v\n", str)
	// better bytes to string
	str = *(*string)(unsafe.Pointer(&bytesRead))
	fmt.Printf("str = %+v\n", str)
}
