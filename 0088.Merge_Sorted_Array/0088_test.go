package leetcode

import "fmt"

func Example_case1() {
	nums := []int{1, 2, 3, 0, 0, 0}
	merge(nums, 3, []int{2, 5, 6}, 3)
	fmt.Print(nums)
	// Output:
	// [1 2 2 3 5 6]
}
