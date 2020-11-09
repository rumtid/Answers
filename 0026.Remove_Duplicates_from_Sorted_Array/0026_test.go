package leetcode

import "fmt"

func Example_case1() {
	nums := []int{1, 1, 2}
	len := removeDuplicates(nums)
	fmt.Print(nums[:len])
	// Output:
	// [1 2]
}

func Example_case2() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len := removeDuplicates(nums)
	fmt.Print(nums[:len])
	// Output:
	// [0 1 2 3 4]
}
