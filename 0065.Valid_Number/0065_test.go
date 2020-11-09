package leetcode

import "fmt"

func Example_case1() {
	ans := isNumber("0")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := isNumber(" 0.1 ")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case3() {
	ans := isNumber("abc")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case4() {
	ans := isNumber("1 a")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case5() {
	ans := isNumber("2e10")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case6() {
	ans := isNumber(" -90e3   ")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case7() {
	ans := isNumber(" 1e")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case8() {
	ans := isNumber("e3")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case9() {
	ans := isNumber(" 6e-1")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case10() {
	ans := isNumber(" 99e2.5 ")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case11() {
	ans := isNumber("53.5e93")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case12() {
	ans := isNumber(" --6 ")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case13() {
	ans := isNumber("-+3")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case14() {
	ans := isNumber("95a54e53")
	fmt.Print(ans)
	// Output:
	// false
}
