package leetcode

import "fmt"

func Example_case1() {
	ans := myAtoi("42")
	fmt.Print(ans)
	// Output:
	// 42
}

func Example_case2() {
	ans := myAtoi("   -42")
	fmt.Print(ans)
	// Output:
	// -42
}

func Example_case3() {
	ans := myAtoi("4193 with words")
	fmt.Print(ans)
	// Output:
	// 4193
}

func Example_case4() {
	ans := myAtoi("words and 987")
	fmt.Print(ans)
	// Output:
	// 0
}

func Example_case5() {
	ans := myAtoi("-91283472332")
	fmt.Print(ans)
	// Output:
	// -2147483648
}
