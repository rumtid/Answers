package main

import "fmt"

func Example_case1() {
	nums := []int{1, 1, 1, 2, 2, 3}
	n := removeDuplicates(nums)
	fmt.Print(nums[:n])
	// Output:
	// [1 1 2 2 3]
}

func Example_case2() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	n := removeDuplicates(nums)
	fmt.Print(nums[:n])
	// Output:
	// [0 0 1 1 2 3 3]
}
