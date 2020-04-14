package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	vals, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, root.Val)
		root = root.Right
	}
	return vals
}
