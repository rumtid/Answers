package main

import "fmt"

func Example_case1() {
	ans := generateTrees(3)
	for _, root := range ans {
		fmt.Println(traverse(root))
	}
	// Unordered output:
	// [1 [] [3 [2] []]]
	// [3 [2 [1] []] []]
	// [3 [1 [] [2]] []]
	// [2 [1] [3]]
	// [1 [] [2 [] [3]]]
}

func traverse(root *TreeNode) (tree []interface{}) {
	if root != nil {
		tree = append(tree, root.Val)
		if root.Left != nil || root.Right != nil {
			tree = append(tree, traverse(root.Left))
			tree = append(tree, traverse(root.Right))
		}
	}
	return
}
