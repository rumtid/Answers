package main

import "fmt"

func Example_case1() {
	ans := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Print(ans)
	// Output:
	// 4
}

func Example_case2() {
	ans := search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Print(ans)
	// Output:
	// -1
}
