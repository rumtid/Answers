package main

import "fmt"

func Example_case1() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Print(nums)
	// Output:
	// [0 0 1 1 2 2]
}
