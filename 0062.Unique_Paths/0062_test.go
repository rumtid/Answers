package leetcode

import "fmt"

func Example_case1() {
	ans := uniquePaths(3, 2)
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := uniquePaths(7, 3)
	fmt.Print(ans)
	// Output:
	// 28
}
