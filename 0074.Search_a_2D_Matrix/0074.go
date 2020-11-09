package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix[0]), len(matrix)
	for lo, hi := 0, m*n-1; lo <= hi; {
		mi := (lo + hi) / 2
		if matrix[mi/m][mi%m] > target {
			hi = mi - 1
		} else if matrix[mi/m][mi%m] < target {
			lo = mi + 1
		} else {
			return true
		}
	}
	return false
}
