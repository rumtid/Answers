package leetcode

func grayCode(n int) []int {
	nums := make([]int, 1<<n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i ^ (i >> 1)
	}
	return nums
}
