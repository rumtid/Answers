package main

import "fmt"

func Example_case1() {
	ans := reverse(123)
	fmt.Print(ans)
	// Output:
	// 321
}

func Example_case2() {
	ans := reverse(-123)
	fmt.Print(ans)
	// Output:
	// -321
}

func Example_case3() {
	ans := reverse(120)
	fmt.Print(ans)
	// Output:
	// 21
}
