package main

import "fmt"

func main() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Printf("a0 = %+v\n", a)
	}

	a = 1
	for a <= 3 {
		a++
		fmt.Printf("a1 = %+v\n", a)
	}

	for a := 1; a < 3; {
		a++
		fmt.Printf("a2 = %+v\n", a)
	}

	a = 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Printf("a3 = %+v\n", a)
	}
}
