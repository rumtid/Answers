package leetcode

func setZeroes(matrix [][]int) {
	m, n, z := len(matrix[0]), len(matrix), false
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			z = true
		}
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if z {
			matrix[i][0] = 0
		}
	}
}
