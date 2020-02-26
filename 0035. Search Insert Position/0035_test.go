package main

import "fmt"

func Example_case1() {
	ans := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case2() {
	ans := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Print(ans)
	// Output:
	// 1
}

func Example_case3() {
	ans := searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Print(ans)
	// Output:
	// 4
}

func Example_case4() {
	ans := searchInsert([]int{1, 3, 5, 6}, 0)
	fmt.Print(ans)
	// Output:
	// 0
}
