package main

import "fmt"

func main() {
	a := 10
	if a > 0 {
		fmt.Printf("a = %+v\n", a)
	}

	if a := 2; a > 1 {
		fmt.Printf("a = %+v\n", a)
	}

	fmt.Printf("a = %+v\n", a)
}
