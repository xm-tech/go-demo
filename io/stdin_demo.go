package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %+v\n", input)
}
