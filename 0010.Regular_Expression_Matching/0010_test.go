package leetcode

import "fmt"

func Example_case1() {
	ans := isMatch("aa", "a")
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case2() {
	ans := isMatch("aa", "a*")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case3() {
	ans := isMatch("ab", ".*")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case4() {
	ans := isMatch("aab", "c*a*b")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case5() {
	ans := isMatch("mississippi", "mis*is*p*.")
	fmt.Print(ans)
	// Output:
	// false
}
