package main

import "fmt"

func main() {
	var b [2]int
	b[1] = 2
	fmt.Printf("b = %+v\n", b)

	// int slice 默认是0，可具体指定每个索引编号对应的值
	a := [...]int{0: 1, 1: 2, 3: 3}
	fmt.Printf("a = %+v\n", a)

	c := [...]int{99: 5}
	fmt.Printf("c = %+v\n", c)

	// 指向数组的指针
	var p *[100]int = &c
	fmt.Printf("p = %+v\n", p)

	// 指针数组
	x, y := 1, 2
	d := [...]*int{&x, &y}
	fmt.Printf("d = %+v\n", d)

	// e := [2]int{1, 2}
	// f := [1]int{2}
	// fmt.Printf("e == f %+v\n", e == f)

	// new 默认也为指向数组的指针
	z := new([10]int)
	z[1] = 2
	fmt.Printf("z = %+v\n", z)

	h := [2][3]int{
		{1, 1, 1},
		{3, 3, 3},
	}
	fmt.Printf("h = %+v\n", h)

	i := [2][3]int{
		{1: 1},
		{2: 3},
	}
	i[0][0] = 5
	fmt.Printf("i = %+v\n", i)

	bubbleSort()
}

// 冒泡排序
func bubbleSort() {
	a := [...]int{5, 3, 1, 8, 2, 9}
	fmt.Printf("a = %+v\n", a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Printf("a = %+v\n", a)
}
