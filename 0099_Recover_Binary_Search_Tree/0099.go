package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var stack []*TreeNode
	var prev, node1, node2 *TreeNode
	for root != nil || len(stack) != 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > root.Val {
			if node1 == nil {
				node1 = prev
			}
			node2 = root
		}
		prev, root = root, root.Right
	}
	node1.Val, node2.Val = node2.Val, node1.Val
}
