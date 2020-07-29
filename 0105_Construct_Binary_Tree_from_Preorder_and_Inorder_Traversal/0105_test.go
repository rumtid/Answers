package main

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := buildTree(
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
	)
	fmt.Print(utils.Ntoa(root))
	// Output:
	// [3 [9] [20 [15] [7]]]
}
