package leetcode

func uniquePaths(m int, n int) int {
	buf := make([]int, m)
	buf[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			buf[j] += buf[j-1]
		}
	}
	return buf[m-1]
}
