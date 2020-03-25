package main

import "fmt"

func Example_case1() {
	ans := search([]int{2, 5, 6, 0, 0, 1, 2}, 0)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := search([]int{2, 5, 6, 0, 0, 1, 2}, 3)
	fmt.Print(ans)
	// Output:
	// false
}
