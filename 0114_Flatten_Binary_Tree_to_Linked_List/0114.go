package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for ; root != nil; root = root.Right {
		if root.Left == nil {
			continue
		}
		node := root.Left
		for ; node.Right != nil; node = node.Right {
		}
		node.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
