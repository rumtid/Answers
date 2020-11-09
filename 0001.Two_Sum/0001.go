package leetcode

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, n := range nums {
		if j, ok := hash[n]; ok {
			return []int{j, i}
		}
		hash[target-n] = i
	}
	return []int{}
}
