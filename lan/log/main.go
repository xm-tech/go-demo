package main

import "log"

func main() {
	log.Println("Hello Golang")
	log.Printf("Hello %+v", "Vim")

	// will exit after exeute
	log.Fatalf("hello %+v", "go")

	log.Println("hello go~")
}
