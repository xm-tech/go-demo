package main

import "fmt"

type Cat struct {
	name string
	age  int
	sex  int
}

func (cat *Cat) Name() {
	fmt.Printf("name = %+v\n", cat.name)
}

func (cat *Cat) Age() {
	fmt.Printf("age = %+v\n", cat.age)
}

func (cat *Cat) Sex() {
	fmt.Printf("sex = %+v\n", cat.sex)
}

func main() {
	// slice
	a := []int{1, 2, 3}
	a[1] = 9
	a = append(a, 4)
	fmt.Println("a=", a)

	// map
	numbers := make(map[string]int, 3)
	numbers["age"] = 1
	fmt.Println(numbers)

	for _, v := range numbers {
		fmt.Println("v=", v)
	}

	// pointer
	n := 0
	add(&n)
	fmt.Println("n=", n)

	// struct demo
	newCat := &Cat{name: "kitty", age: 3, sex: 0}
	newCat.age = 4
	fmt.Printf("newCat = %+v\n", newCat)
	newCat.Name()
	newCat.Age()
	newCat.Sex()

}

func add(n *int) {
	*n = *n + 1
}
