package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int
	queue := []*TreeNode{root}
	for {
		var level []int
		n := len(queue)
		for i := 0; i < n; i++ {
			if root = queue[i]; root != nil {
				level = append(level, root.Val)
				queue = append(queue, root.Left)
				queue = append(queue, root.Right)
			}
		}
		if len(level) == 0 {
			break
		}
		levels, queue = append(levels, level), queue[n:]
	}
	return levels
}
