package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	t, buf := 0, make([]int, m+1)
	for i := 0; i <= m; i++ {
		buf[i] = i
	}
	for i := 1; i <= n; i++ {
		t, buf[0] = buf[0], buf[0]+1
		for j := 1; j <= m; j++ {
			if word1[j-1] != word2[i-1] {
				t++
			}
			t, buf[j] = buf[j], min(buf[j]+1, buf[j-1]+1, t)
		}
	}
	return buf[m]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}
