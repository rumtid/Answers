package leetcode

func removeDuplicates(nums []int) int {
	i, j := 2, 2
	for ; j < len(nums); j++ {
		if nums[j] != nums[i-2] {
			nums[i], i = nums[j], i+1
		}
	}
	return i
}
