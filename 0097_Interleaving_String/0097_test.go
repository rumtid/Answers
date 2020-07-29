package main

import "fmt"

func Example_case1() {
	ans := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := isInterleave("aabcc", "dbbca", "aadbbbaccc")
	fmt.Print(ans)
	// Output:
	// false
}
