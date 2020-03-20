package main

var set []int
var sets [][]int

func subsets(nums []int) [][]int {
	sets = [][]int{}
	for i := 0; i <= len(nums); i++ {
		search(nums, i)
	}
	return sets
}

func search(nums []int, k int) {
	if k == 0 {
		sets = append(sets, append([]int{}, set...))
		return
	}
	for i := 0; i < len(nums)-k+1; i++ {
		set = append(set, nums[i])
		search(nums[i+1:], k-1)
		set = set[:len(set)-1]
	}
}
