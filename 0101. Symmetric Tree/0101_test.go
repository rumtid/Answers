package main

import (
	"../utils"
	"fmt"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[1 [2 [3] [4]] [2 [4] [3]]]"))
	ans := isSymmetric(root)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	root := (*TreeNode)(utils.Aton("[1 [2 [] [3]] [2 [] [3]]]"))
	ans := isSymmetric(root)
	fmt.Print(ans)
	// Output:
	// false
}
