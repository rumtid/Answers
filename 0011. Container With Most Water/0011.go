package main

func maxArea(height []int) int {
	max := 0
	for lo, hi := 0, len(height)-1; lo < hi; {
		if height[lo] < height[hi] {
			if max < height[lo]*(hi-lo) {
				max = height[lo] * (hi - lo)
			}
			lo++
		} else {
			if max < height[hi]*(hi-lo) {
				max = height[hi] * (hi - lo)
			}
			hi--
		}
	}
	return max
}
