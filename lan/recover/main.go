package main

import (
	"fmt"
	"log"
)

func pcall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("f panic and can not recover from err:%v", err)
		}
	}()

	f()
	fmt.Println("after f")
}

func main() {
	e := func() {
		panic("an runtime err")
	}
	// e()
	pcall(e)

	log.Println("assume will be execute normally")
}
