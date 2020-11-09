package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			if i+1 < j {
				nums[i+1] = nums[j]
			}
			i++
		}
		j++
	}
	return i + 1
}
