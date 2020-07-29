package main

import "fmt"

func Example_case1() {
	ans := isScramble("great", "rgeat")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := isScramble("abcde", "caebd")
	fmt.Print(ans)
	// Output:
	// false
}
