package main

func trap(height []int) int {
	var level, water int
	for lo, hi := 0, len(height)-1; lo < hi; {
		for ; lo < hi && height[lo] <= level; lo++ {
			water += level - height[lo]
		}
		for ; lo < hi && height[hi] <= level; hi-- {
			water += level - height[hi]
		}
		if height[lo] < height[hi] {
			level = height[lo]
		} else {
			level = height[hi]
		}
	}
	return water
}
