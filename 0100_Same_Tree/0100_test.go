package main

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	p := (*TreeNode)(utils.Aton("[1 [2] [3]]"))
	q := (*TreeNode)(utils.Aton("[1 [2] [3]]"))
	ans := isSameTree(p, q)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	p := (*TreeNode)(utils.Aton("[1 [2] []]"))
	q := (*TreeNode)(utils.Aton("[1 [] [2]]"))
	ans := isSameTree(p, q)
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case3() {
	p := (*TreeNode)(utils.Aton("[1 [2] [1]]"))
	q := (*TreeNode)(utils.Aton("[1 [1] [2]]"))
	ans := isSameTree(p, q)
	fmt.Print(ans)
	// Output:
	// false
}
