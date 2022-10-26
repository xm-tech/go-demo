package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// pprof
	go func() {
		http.ListenAndServe("0.0.0.0:10000", nil)
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 50000; i++ {
			fmt.Printf("A%+v = %+v\n", i, i)
		}
	}()

	fmt.Println("===========================")

	go func() {
		defer wg.Done()
		for i := 1; i < 50000; i++ {
			fmt.Printf("B%+v = %+v\n", i, i)
		}
	}()

	wg.Wait()

	for {

	}
}
