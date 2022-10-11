package main

import (
	"log"
	"time"
)

func main() {
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
