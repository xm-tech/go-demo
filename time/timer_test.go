package main

import (
	"log"
	"testing"
	"time"
)

func setup() {

}

func teardown() {

}

func TestAfter(t *testing.T) {
	// t.Helper()
	// log.Println("begin")

	// go time.AfterFunc(3*time.Second, func() {
	// 	log.Println("3 sec after ...")
	// })

	// // <-time.After(3 * time.Second)

	// log.Println("3 sec after")

	// select {}
}

func TestChanAfter(t *testing.T) {

	log.Println("begin")

	status := "cmping"

	timer := time.AfterFunc(3*time.Second, func() {
		status = "over"
		log.Println("status:", status)
	})

	log.Println("........,", timer)

	// timer.Stop()

	select {}
}
