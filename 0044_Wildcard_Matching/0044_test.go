package main

import "fmt"

func Example_case1() {
	ans := isMatch("aa", "a")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case2() {
	ans := isMatch("aa", "*")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case3() {
	ans := isMatch("cb", "?a")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case4() {
	ans := isMatch("adceb", "*a*b")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case5() {
	ans := isMatch("acdcb", "a*c?b")
	fmt.Print(ans)
	// Output:
	// false
}
