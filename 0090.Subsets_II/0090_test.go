package leetcode

import "fmt"

func Example_case1() {
	ans := subsetsWithDup([]int{1, 2, 2})
	for _, set := range ans {
		fmt.Println(set)
	}
	// Unordered output:
	// [2]
	// [1]
	// [1 2 2]
	// [2 2]
	// [1 2]
	// []
}
