package main

func minPathSum(grid [][]int) int {
	m, n := len(grid[0]), len(grid)
	buf := make([]int, m)
	for i := 0; i < n; i++ {
		buf[0] += grid[i][0]
		for j := 1; j < m; j++ {
			if i == 0 || buf[j-1] < buf[j] {
				buf[j] = buf[j-1]
			}
			buf[j] += grid[i][j]
		}
	}
	return buf[m-1]
}
