package main

import "fmt"

func Example_case1() {
	ans := divide(10, 3)
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := divide(7, -3)
	fmt.Print(ans)
	// Output:
	// -2
}
