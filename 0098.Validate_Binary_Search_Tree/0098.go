package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev, root = root, root.Right
	}
	return true
}
