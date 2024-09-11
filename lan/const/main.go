package main

import (
	"fmt"
	"log"
)

const (
	MON = iota
	TUS
	WEN
	THR
	FRI
	SAT
	SUN
)

const (
	A = 1
	// iota , 常量组中每定义1个常量，iota 递增1, 初值是 0
	B = iota
	// C 的值是b的值
	C
)

const (
	_D = iota
)

type ByteSize float64

func (self ByteSize) info() {
	fmt.Printf("val= %+v\n", self)
}

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("MON = %+v\n", MON)
	fmt.Printf("TUS = %+v\n", TUS)
	fmt.Printf("WEN = %+v\n", WEN)
	fmt.Printf("THR = %+v\n", THR)
	fmt.Printf("FRI = %+v\n", FRI)
	fmt.Printf("SAT = %+v\n", SAT)
	fmt.Printf("SUN = %+v\n", SUN)

	fmt.Printf("A = %+v\n", A)
	fmt.Printf("B = %+v\n", B)
	fmt.Printf("C = %+v\n", C)

	fmt.Printf("_D = %+v\n", _D)

	KB.info()
	MB.info()
	GB.info()
	TB.info()
	PB.info()
	EB.info()
	ZB.info()
	YB.info()
}

func ptrOp() {
	// 常量赋值给指针

	type Man struct {
		Age int
	}

	man := &Man{
		Age: A,
	}

	log.Println("man:", man)
}
