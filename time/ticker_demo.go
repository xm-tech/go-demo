package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// testTicker()
	f := func() {
		fmt.Println("f exec")
	}
	// ThenDo(3*time.Second, f)

	time.AfterFunc(3*time.Second, f)
	for {

	}
}

func testTicker() {
	done := make(chan struct{}, 1)
	defer func() {
		close(done)
	}()

	tick := time.NewTicker(time.Second * 1)
	go func() {
		loop := 0
		for {
			<-tick.C
			loop++
			log.Println("dida ", loop)
			if loop >= 5 {
				tick.Stop()
				done <- struct{}{}
				break
			}
		}
	}()

	log.Println("wait done")
	<-done
	log.Println("done")
}

// exec funciton f after sec second
func ThenDo(duration time.Duration, f func()) {
	fmt.Println("ThenDo Begin", time.Now().UTC())
	<-time.After(duration)
	fmt.Println("ThenDo End, ", time.Now().UTC())
	f()
}
