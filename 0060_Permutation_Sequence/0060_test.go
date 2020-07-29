package main

import "fmt"

func Example_case1() {
	ans := getPermutation(3, 3)
	fmt.Print(ans)
	// Output:
	// 213
}

func Example_case2() {
	ans := getPermutation(4, 9)
	fmt.Print(ans)
	// Output:
	// 2314
}
