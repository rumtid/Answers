package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return traverse(root.Left, root.Right)
}

func traverse(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	} else if p.Val != q.Val {
		return false
	}
	return traverse(p.Left, q.Right) && traverse(p.Right, q.Left)
}
