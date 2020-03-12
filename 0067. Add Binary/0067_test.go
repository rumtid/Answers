package main

import "fmt"

func Example_case1() {
	ans := addBinary("11", "1")
	fmt.Print(ans)
	// Output:
	// 100
}

func Example_case2() {
	ans := addBinary("1010", "1011")
	fmt.Print(ans)
	// Output:
	// 10101
}
