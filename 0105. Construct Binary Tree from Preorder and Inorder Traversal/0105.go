package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode
	if len(preorder) > 0 {
		root = &TreeNode{Val: preorder[0]}
		i := 0
		for ; inorder[i] != preorder[0]; i++ {
		}
		root.Left = buildTree(preorder[1:i+1], inorder[:i])
		root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}
	return root
}
