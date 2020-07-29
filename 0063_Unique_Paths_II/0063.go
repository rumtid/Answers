package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid[0]), len(obstacleGrid)
	buf := make([]int, m)
	buf[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				buf[j] = 0
			} else if j > 0 {
				buf[j] += buf[j-1]
			}
		}
	}
	return buf[m-1]
}
