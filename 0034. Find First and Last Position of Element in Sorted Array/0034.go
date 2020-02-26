package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	i, j, m, n := 0, 0, -1, -1
	for i, j = 0, len(nums)-1; i < j; {
		if nums[(i+j)/2] < target {
			i = (i+j)/2 + 1
		} else {
			j = (i + j) / 2
		}
	}
	if nums[i] == target {
		m = i
	}
	for i, j = 0, len(nums)-1; i < j; {
		if nums[(i+j+1)/2] > target {
			j = (i+j+1)/2 - 1
		} else {
			i = (i + j + 1) / 2
		}
	}
	if nums[i] == target {
		n = i
	}
	return []int{m, n}
}
