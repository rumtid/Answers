package leetcode

import "fmt"

func Example_case1() {
	ans := isPalindrome(121)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := isPalindrome(-121)
	fmt.Print(ans)
	// Output:
	// false
}

func Example_case3() {
	ans := isPalindrome(10)
	fmt.Print(ans)
	// Output:
	// false
}
