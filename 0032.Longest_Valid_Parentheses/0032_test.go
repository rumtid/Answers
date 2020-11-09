package leetcode

import "fmt"

func Example_case1() {
	ans := longestValidParentheses("(()")
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case2() {
	ans := longestValidParentheses(")()())")
	fmt.Print(ans)
	// Output:
	// 4
}
