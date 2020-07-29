package main

func maximalRectangle(matrix [][]byte) int {
	max, stack, n := 0, []int{}, 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] -= '0'
			if matrix[i][j] == 1 && i > 0 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
		heights := append(matrix[i], 0)
		n = 0
		for j, h := range heights {
			for ; n > 0 && h < heights[stack[n-1]]; n-- {
				w := j
				if n > 1 {
					w -= int(stack[n-2]) + 1
				}
				if int(heights[stack[n-1]])*w > max {
					max = int(heights[stack[n-1]]) * w
				}
			}
			stack, n = append(stack[:n], j), n+1
		}
	}

	return max
}
