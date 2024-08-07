package main

import "fmt"

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Printf("m1 = %+v\n", m1)
	fmt.Printf("m11 = %v\n", m1)

	m2 := map[string]int{}
	// m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Printf("m2 = %+v\n", m2)

	keys := make([]string, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}

	fmt.Printf("keys = %+v\n", keys)
}
