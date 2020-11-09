package leetcode

func numTrees(n int) int {
	buf := make([]int, n+1)
	buf[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			buf[i] += buf[j] * buf[i-j-1]
		}
	}
	return buf[n]
}
