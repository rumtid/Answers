package leetcode

func sortColors(nums []int) {
	lo, hi := 0, len(nums)-1
	for i := 0; i <= hi; {
		switch nums[i] {
		case 0:
			nums[i], nums[lo] = nums[lo], nums[i]
			i, lo = i+1, lo+1
		case 2:
			nums[i], nums[hi] = nums[hi], nums[i]
			hi--
		case 1:
			i++
		}
	}
}
