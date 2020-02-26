package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	quadruplets := [][]int{}
	for i, a := range nums {
		if i != 0 && a == nums[i-1] {
			continue
		}
		for j, b := range nums[i+1:] {
			if j != 0 && b == nums[i+j] {
				continue
			}
			for lo, hi := i+j+2, len(nums)-1; lo < hi; {
				c := nums[lo]
				d := nums[hi]
				s := a + b + c + d
				if s < target {
					lo++
				} else if s > target {
					hi--
				} else {
					quadruplets = append(quadruplets, []int{a, b, c, d})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				}
			}
		}
	}
	return quadruplets
}
