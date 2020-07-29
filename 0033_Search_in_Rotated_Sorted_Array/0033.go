package main

func search(nums []int, target int) int {
	var i, j int
	for i, j = 0, len(nums)-1; i < j; {
		if nums[(i+j)/2] > nums[j] {
			i = (i+j)/2 + 1
		} else {
			j = (i + j) / 2
		}
	}
	d := i
	for i, j = 0, len(nums)-1; i <= j; {
		n := nums[((i+j)/2+d)%len(nums)]
		if n == target {
			return ((i+j)/2 + d) % len(nums)
		} else if n < target {
			i = (i+j)/2 + 1
		} else {
			j = (i+j)/2 - 1
		}
	}
	return -1
}
