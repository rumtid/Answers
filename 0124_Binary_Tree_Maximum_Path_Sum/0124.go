package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func maxPathSum(root *TreeNode) int {
	max = math.MinInt32
	traverse(root)
	return max
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	traverse(root.Right)
	var l, r int
	if root.Left != nil {
		l = root.Left.Val
	}
	if root.Right != nil {
		r = root.Right.Val
	}
	if root.Val+l+r > max {
		max = root.Val + l + r
	}
	if l < r {
		l = r
	}
	if l > 0 {
		root.Val += l
	}
	if root.Val > max {
		max = root.Val
	}
}
