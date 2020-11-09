package leetcode

func canJump(nums []int) bool {
	var j int
	for i, n := range nums {
		if i > j {
			return false
		}
		if i+n > j {
			j = i + n
		}
	}
	return true
}
