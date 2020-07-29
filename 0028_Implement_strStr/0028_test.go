package main

import "fmt"

func Example_case1() {
	ans := strStr("hello", "ll")
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case2() {
	ans := strStr("aaaaa", "bba")
	fmt.Print(ans)
	// Output:
	// -1
}
