package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	// f()
	pcall(f)

	for {
		time.Sleep(1 * time.Second)
	}
}

func f() {
	defer func() {
		fmt.Println("f end")
	}()
	panic("nil")
}

func pcall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err captured: ", err)
			debug.PrintStack()
		}
	}()
	f()
}
