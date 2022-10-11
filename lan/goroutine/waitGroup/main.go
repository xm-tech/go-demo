package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Printf("A%+v = %+v\n", i, i)
		}
	}()

	fmt.Println("===========================")

	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Printf("B%+v = %+v\n", i, i)
		}
	}()

	wg.Wait()
}
