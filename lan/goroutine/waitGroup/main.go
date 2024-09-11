package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	// // runtime.GOMAXPROCS(1)
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// // pprof
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:10000", nil)
	// }()

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i < 50000; i++ {
	// 		fmt.Printf("A%+v = %+v\n", i, i)
	// 	}
	// }()

	// fmt.Println("===========================")

	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i < 50000; i++ {
	// 		fmt.Printf("B%+v = %+v\n", i, i)
	// 	}
	// }()

	// wg.Wait()

	// log.Println("end")

	// for {

	// }

	f()
}

var exitSync *sync.WaitGroup
var closing int64

func f() {

	exitSync = &sync.WaitGroup{}
	exitSync.Add(3)
	atomic.StoreInt64(&closing, 0)

	go a()
	go b()
	go c()

	Close()

	exitSync.Wait()
	log.Println("end")
}

func a() {
	for {
		if atomic.LoadInt64(&closing) == 1 {
			break
		}
	}
	exitSync.Done()
}

func b() {
	for {
		if atomic.LoadInt64(&closing) == 1 {
			break
		}
	}
	exitSync.Done()
}

func c() {
	for {
		if atomic.LoadInt64(&closing) == 1 {
			break
		}
	}
	exitSync.Done()
}

func Close() {
	atomic.StoreInt64(&closing, 1)
}
