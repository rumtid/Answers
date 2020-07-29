package main

import "fmt"

func Example_case1() {
	ans := numDecodings("12")
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case2() {
	ans := numDecodings("226")
	fmt.Print(ans)
	// Output:
	// 3
}
