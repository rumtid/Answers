package main

import "sort"

var set []int
var sets [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	sets = [][]int{}
	for i := 0; i <= len(nums); i++ {
		search(nums, i)
	}
	return sets
}

func search(nums []int, n int) {
	if n == 0 {
		sets = append(sets, append([]int{}, set...))
		return
	}
	for i := 0; i < len(nums)-n+1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		set = append(set, nums[i])
		search(nums[i+1:], n-1)
		set = set[:len(set)-1]
	}
}
