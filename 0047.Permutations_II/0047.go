package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	permutations := [][]int{}
	permutations = append(permutations, append([]int{}, nums...))
	for nextPermutation(nums) {
		permutations = append(permutations, append([]int{}, nums...))
	}
	return permutations
}

func nextPermutation(nums []int) bool {
	i, n := 0, len(nums)
	for i = n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i == 0 {
		return false
	}
	for j := n - 1; j > i-1; j-- {
		if nums[j] > nums[i-1] {
			nums[j], nums[i-1] = nums[i-1], nums[j]
			break
		}
	}
	for j := n - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return true
}
