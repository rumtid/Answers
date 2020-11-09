package leetcode

import "fmt"

func Example_case1() {
	ans := isValid("()")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := isValid("()[]{}")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case3() {
	ans := isValid("(]")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case4() {
	ans := isValid("([)]")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case5() {
	ans := isValid("{[]}")
	fmt.Print(ans)
	// Output:
	// true
}
