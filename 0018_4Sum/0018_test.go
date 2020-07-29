package main

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [-1 0 0 1]
	// [-2 -1 1 2]
	// [-2 0 0 2]
}
