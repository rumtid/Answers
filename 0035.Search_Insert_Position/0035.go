package leetcode

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[(i+j)/2] < target {
			i = (i+j)/2 + 1
		} else {
			j = (i + j) / 2
		}
	}
	return i
}
