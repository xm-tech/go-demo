package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 1:
		fmt.Printf("a = %+v\n", a)
	case 2:
		fmt.Printf("a = %+v\n", a)
	default:
		fmt.Printf("a = %+v\n", a)
	}

	switch a := 1; {
	case a == 1:
		fmt.Printf("a = %+v\n", a)
	default:
		fmt.Printf("a = %+v\n", a)
	}

	a = 1
	// 当有1个case 条件表达式满足时,switch 终止
	switch {
	case a >= 0:
		fmt.Printf("a >= %+v\n", 0)
		// 不退出, switch 继续进行下一条 case
		fallthrough
	case a >= 1:
		fmt.Printf("a >= %+v\n", 1)
	default:
		fmt.Printf("None \n ")
	}
}
