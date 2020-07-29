package main

func largestRectangleArea(heights []int) int {
	max, stack, n := 0, []int{}, 0
	heights = append(heights, 0)
	for i, h := range heights {
		for ; n > 0 && h < heights[stack[n-1]]; n-- {
			w := i
			if n > 1 {
				w -= stack[n-2] + 1
			}
			if heights[stack[n-1]]*w > max {
				max = heights[stack[n-1]] * w
			}
		}
		stack, n = append(stack[:n], i), n+1
	}
	return max
}
