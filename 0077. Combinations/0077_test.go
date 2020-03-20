package main

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := combine(4, 2)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [2 4]
	// [3 4]
	// [2 3]
	// [1 2]
	// [1 3]
	// [1 4]
}
