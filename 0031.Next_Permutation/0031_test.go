package leetcode

import "fmt"

func Example_case1() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Print(nums)
	// Output:
	// [1 3 2]
}

func Example_case2() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Print(nums)
	// Output:
	// [1 2 3]
}

func Example_case3() {
	nums := []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Print(nums)
	// Output:
	// [1 5 1]
}
