package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	} else {
		return generate(0, n, []byte{})
	}
}

func generate(m, n int, s []byte) []string {
	if m == 0 && n == 0 {
		return []string{string(s)}
	}
	strs := make([]string, 0)
	if n > 0 {
		strs = append(strs, generate(m+1, n-1, append(s, '('))...)
	}
	if m > 0 {
		strs = append(strs, generate(m-1, n, append(s, ')'))...)
	}
	return strs
}
