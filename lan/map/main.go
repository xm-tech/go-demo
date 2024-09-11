package main

import (
	"fmt"
	"log"
)

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

	testNil()
}

func testNil() {
	uidGunUids := make(map[int64][]int64)

	uids := []int64{100, 101, 102, 103}
	for _, v := range uids {
		if v != 100 {
			uidGunUids[100] = append(uidGunUids[100], v)
		}
	}

	log.Printf("uidGunUids:%+v", uidGunUids)
}
