package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var s, t []*TreeNode
	s = append(s, root)
	for len(s) > 0 {
		var a []int
		for ; len(s) > 0; s = s[1:] {
			if s[0].Left != nil {
				t = append(t, s[0].Left)
			}
			if s[0].Right != nil {
				t = append(t, s[0].Right)
			}
			a = append(a, s[0].Val)
		}
		ans = append(ans, a)
		s, t = t, nil
	}
	for i := 1; i < len(ans); i += 2 {
		for j, k := 0, len(ans[i])-1; j < k; j, k = j+1, k-1 {
			ans[i][j], ans[i][k] = ans[i][k], ans[i][j]
		}
	}
	return ans
}
