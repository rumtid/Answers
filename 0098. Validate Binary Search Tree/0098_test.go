package main

import (
	"../utils"
	"fmt"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[2 [1] [3]]"))
	ans := isValidBST(root)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	root := (*TreeNode)(utils.Aton("[5 [1] [4 [3] [6]]]"))
	ans := isValidBST(root)
	fmt.Print(ans)
	// Output:
	// false
}
