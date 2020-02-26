package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var i, j, k int
	sum := nums[0] + nums[1] + nums[2]
	for i = 0; i < len(nums)-2; i++ {
		j = i + 1
		k = len(nums) - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if abs(s-target) < abs(sum-target) {
				sum = s
			}
			if s-target > 0 {
				k--
			} else if s-target < 0 {
				j++
			} else {
				return s
			}
		}
	}
	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
