package leetcode

func longestValidParentheses(s string) int {
	var l, r, max int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
			continue
		}
		r++
		if l == r && l+r > max {
			max = l + r
		} else if l < r {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			r++
			continue
		}
		l++
		if l == r && l+r > max {
			max = l + r
		} else if l > r {
			l, r = 0, 0
		}
	}
	return max
}
