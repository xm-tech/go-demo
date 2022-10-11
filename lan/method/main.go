package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	A
	Name string
}

type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}

func main() {
	b := B{A: A{Name: "A"}, Name: "B"}
	a := A{Name: "A"}

	fmt.Printf("b = %+v\n", b)
	fmt.Printf("a = %+v\n", a)

	var c TZ = 0
	c.Increase(150)
	fmt.Printf("c = %+v\n", c)
}
