package main

import "fmt"

func Example_case1() {
	ans := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Print(ans)
	// Output:
	// [3 4]
}

func Example_case2() {
	ans := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Print(ans)
	// Output:
	// [-1 -1]
}
