package main

import (
	"fmt"
	"strconv"
)

func Example_case1() {
	root, _ := construct("[2 [1] [3]]")
	ans := isValidBST(root)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	root, _ := construct("[5 [1] [4 [3] [6]]]")
	ans := isValidBST(root)
	fmt.Print(ans)
	// Output:
	// false
}

func construct(s string) (*TreeNode, string) {
	if s[0] != '[' {
		panic("missing '['")
	}
	if s[1] == ']' {
		return nil, s[2:]
	}
	i := 1
	for ; s[i] >= '0' && s[i] <= '9'; i++ {
	}
	root := new(TreeNode)
	root.Val, _ = strconv.Atoi(s[1:i])
	if s = s[i:]; s[0] == ' ' {
		root.Left, s = construct(s[1:])
		root.Right, s = construct(s[1:])
	}
	return root, s[1:]
}
