package main

import "fmt"

func Example_case1() {
	ans := minDistance("horse", "ros")
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := minDistance("intention", "execution")
	fmt.Print(ans)
	// Output:
	// 5
}
