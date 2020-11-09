package leetcode

func lengthOfLastWord(s string) int {
	var i, l int = len(s) - 1, 0
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		l++
	}
	return l
}
