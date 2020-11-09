package leetcode

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := combinationSum([]int{2, 3, 6, 7}, 7)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [7]
	// [2 2 3]
}

func Example_case2() {
	ans := combinationSum([]int{2, 3, 5}, 8)
	for _, set := range ans {
		sort.Ints(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [2 2 2 2]
	// [2 3 3]
	// [3 5]
}
