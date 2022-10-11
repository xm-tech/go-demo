package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("kafka.conf")
	if err != nil {
		panic(err)
	}
	kafkaAddr := string(data)
	fmt.Printf("kafka.addr=%+v", kafkaAddr)
}
