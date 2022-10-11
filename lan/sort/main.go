package main

import (
	"fmt"
	"sort"
)

type StuAge struct {
	Name string
	Age  int
}

type StuAges []StuAge

func (sa StuAges) Len() int {
	return len(sa)
}

func (sa StuAges) Less(i, j int) bool {
	return sa[i].Age < sa[j].Age
}

func (sa StuAges) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func main() {
	stus := StuAges{
		{"maxm", 36},
		{"lfj", 35},
		{"cfm", 33},
	}

	fmt.Printf("stus = %+v\n", stus)
	sort.Sort(stus)
	fmt.Printf("stus = %+v\n", stus)

	sort.Reverse(stus)
	fmt.Printf("stus = %+v\n", stus)
}
