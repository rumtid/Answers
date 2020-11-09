package leetcode

func firstMissingPositive(nums []int) int {
	for _, n := range nums {
		for n > 0 && n <= len(nums) && nums[n-1] != n {
			nums[n-1], n = n, nums[n-1]
		}
	}
	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
