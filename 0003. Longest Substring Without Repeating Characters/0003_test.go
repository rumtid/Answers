package main

import "fmt"

func Example_case1() {
	ans := lengthOfLongestSubstring("abcabcbb")
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := lengthOfLongestSubstring("bbbbb")
	fmt.Print(ans)
	// Output:
	// 1
}

func Example_case3() {
	ans := lengthOfLongestSubstring("pwwkew")
	fmt.Print(ans)
	// Output:
	// 3
}
