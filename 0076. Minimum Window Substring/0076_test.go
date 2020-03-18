package main

import "fmt"

func Example_case1() {
	ans := minWindow("ADOBECODEBANC", "ABC")
	fmt.Print(ans)
	// Output:
	// BANC
}
