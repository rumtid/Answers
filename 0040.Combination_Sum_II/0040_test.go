package leetcode

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [1 7]
	// [1 2 5]
	// [2 6]
	// [1 1 6]
}

func Example_case2() {
	ans := combinationSum2([]int{2, 5, 2, 1, 2}, 5)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [1 2 2]
	// [5]
}
