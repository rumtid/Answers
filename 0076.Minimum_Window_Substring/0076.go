package leetcode

func minWindow(s string, t string) string {
	var buf [128]int
	for i := 0; i < len(t); i++ {
		buf[t[i]]++
	}
	min, num := "", len(t)
	for i, j := 0, 0; j < len(s); j++ {
		if buf[s[j]]--; buf[s[j]] >= 0 {
			num--
		}
		for ; num == 0; i++ {
			if buf[s[i]]++; buf[s[i]] == 1 {
				if len(min) == 0 || j-i+1 < len(min) {
					min = s[i : j+1]
				}
				num++
			}
		}
	}
	return min
}
