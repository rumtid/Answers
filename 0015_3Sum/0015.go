package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}
	for i, a := range nums {
		if i != 0 && a == nums[i-1] {
			continue
		}
		for lo, hi := i+1, len(nums)-1; lo < hi; {
			b := nums[lo]
			c := nums[hi]
			s := a + b + c
			if s < 0 {
				lo++
			} else if s > 0 {
				hi--
			} else {
				triplets = append(triplets, []int{a, b, c})
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
	return triplets
}
