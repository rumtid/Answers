package leetcode

import "fmt"

func Example_case1() {
	ans := subsets([]int{1, 2, 3})
	for _, set := range ans {
		fmt.Println(set)
	}
	// Unordered output:
	// [3]
	// [1]
	// [2]
	// [1 2 3]
	// [1 3]
	// [2 3]
	// [1 2]
	// []
}
