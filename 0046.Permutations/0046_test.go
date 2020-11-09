package leetcode

import "fmt"

func Example_case1() {
	ans := permute([]int{1, 2, 3})
	for _, set := range ans {
		fmt.Println(set)
	}
	// Unordered output:
	// [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 1 2]
	// [3 2 1]
}
