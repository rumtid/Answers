package main

import (
	"../utils"
	"fmt"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[1 [2 [3] [4]] [5 [] [6]]]"))
	flatten(root)
	fmt.Print(utils.Ntoa(root))
	// Output:
	// [1 [] [2 [] [3 [] [4 [] [5 [] [6]]]]]]
}
