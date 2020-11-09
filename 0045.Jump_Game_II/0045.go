package leetcode

func jump(nums []int) int {
	var j, k, s int
	for i, n := range nums {
		if i > j {
			j = k
			s++
		}
		if i+n > k {
			k = i + n
		}
	}
	return s
}
