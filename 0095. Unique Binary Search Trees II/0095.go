package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(lo, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}
	var roots []*TreeNode
	for i := lo; i <= hi; i++ {
		for _, left := range generate(lo, i-1) {
			for _, right := range generate(i+1, hi) {
				roots = append(roots, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return roots
}
