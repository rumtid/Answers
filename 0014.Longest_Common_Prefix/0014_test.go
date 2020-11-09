package leetcode

import "fmt"

func Example_case1() {
	ans := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Print(ans)
	// Output:
	// fl
}

func Example_case2() {
	ans := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Print(ans)
	// Output:
}
