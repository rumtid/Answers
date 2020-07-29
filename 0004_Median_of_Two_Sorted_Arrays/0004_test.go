package main

import "fmt"

func Example_case1() {
	ans := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case2() {
	ans := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Print(ans)
	// Output:
	// 2.5
}
