package leetcode

import "fmt"

func Example_case1() {
	ans := longestPalindrome("babad")
	fmt.Print(ans)
	// Output:
	// aba
}

func Example_case2() {
	ans := longestPalindrome("cbbd")
	fmt.Print(ans)
	// Output:
	// bb
}
