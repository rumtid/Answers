package main

import "fmt"

func Example_case1() {
	ans := multiply("2", "3")
	fmt.Print(ans)
	// Output:
	// 6
}

func Example_case2() {
	ans := multiply("123", "456")
	fmt.Print(ans)
	// Output:
	// 56088
}
