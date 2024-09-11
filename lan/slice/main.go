package main

import (
	"fmt"
	"log"
	"reflect"
)

/*
*
- array
  - 固定长度数组，使用前必须指定数组长度
  - golang中的数组是值类型,也就是说，如果你将一个数组赋值给另外一个数组，那么，实际上就是整个数组拷贝了一份
  - 如果golang中的数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针
  - array的长度也是Type的一部分，这样就说明[10]int和[20]int是不一样的。

- slice
  - slice是一个引用类型，是一个动态的指向数组切片的指针。
  - slice是一个不定长的，总是指向底层的数组array的数据结构。
*/
func main() {
	// f0()
	f1()
}

func f0() {
	// 声明1 slice
	var s1 []int
	fmt.Printf("s1 = %+v\n", s1)
	fmt.Printf("s1 = %+v, s1.type = %T, s1.addr = %p, s1=%+v\n", s1, s1, s1, reflect.TypeOf(s1).String())

	// 声明 1 array
	a := [10]int{}
	fmt.Printf("a = %+v\n", a)

	fmt.Printf("a[9] = %+v\n", a[9])

	// 固定长度数组
	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("b[5:10] = %+v\n", b[5:10])

	// 截取产生 slice
	fmt.Printf("b[5:len(b)] = %+v\n", b[5:len(b)])

	fmt.Printf("b[5:] = %+v\n", b[5:])

	// 另一种 slice 声明方式
	s2 := make([]int, 3, 10)
	fmt.Printf("s2 = %+v,len=%+v,cap=%+v\n", s2, len(s2), cap(s2))

	// 声明1 slice，并初始化
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("s3[:] = %+v\n", s3[:])

	// 切片s4指定长度2(6-4)，指定容量5(9-4)
	s4 := s3[4:6:9]
	fmt.Printf("s4 = %+v,len(s4)=%+v,cap(s4)=%+v\n", s4, len(s4), cap(s4))
	// s5 长度:3-1=2,容量:3-1=2
	s5 := s3[1:3:3]
	fmt.Printf("s5 = %+v,len(s5)=%+v,cap(s5)=%+v\n", s5, len(s5), cap(s5))
	// s6 长度:3-1=2, 容量=底层数组容量-1=5-1=4
	s6 := s3[1:3]
	fmt.Printf("s6 = %+v,len(s6)=%+v,cap(s6)=%+v\n", s6, len(s6), cap(s6))

}

func f1() {

	resp := make([]Msg, 0)
	log.Println("len(resp):", len(resp))
	for _, v := range resp {
		log.Println("---", v)
	}
}

type Msg struct {
}
