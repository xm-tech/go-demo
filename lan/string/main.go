package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str333 := strconv.Itoa(3333)
	fmt.Printf("str333 = %+v\n", str333)

	ret := strings.Compare("abc", "def")
	fmt.Printf("ret = %+v\n", ret)

	contains := strings.Contains("go strings demo", "go1")
	fmt.Printf("contains = %+v\n", contains)

	count := strings.Count("welcome to the golang world", "e")
	fmt.Printf("count = %+v\n", count)

	trim := strings.Trim("  hello go lang  ", " ")
	fmt.Printf("trim = %+v\n", trim)

	trimSpace := strings.TrimSpace("   helllo golang   ")
	fmt.Printf("trimSpace = %+v\n", trimSpace)

	words := strings.Split("maxm,lfj,36", ",")
	for _, v := range words {
		fmt.Printf("%+v\n", v)
	}
	fmt.Printf("words= %+v\n", words)

	hasPrefix := strings.HasPrefix("Friend_List", "Friend")
	fmt.Printf("hasPrefix = %+v\n", hasPrefix)

	hasSuffix := strings.HasSuffix("Friend_List", "t")
	fmt.Printf("hasSuffix = %+v\n", hasSuffix)
}
