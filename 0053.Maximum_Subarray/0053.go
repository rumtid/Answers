package leetcode

func maxSubArray(nums []int) int {
	pre, sum := nums[0], nums[0]
	for _, num := range nums[1:] {
		pre = max(pre+num, num)
		sum = max(sum, pre)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
