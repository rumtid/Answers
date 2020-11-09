package leetcode

import "fmt"

func Example_case1() {
	ans := firstMissingPositive([]int{1, 2, 0})
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := firstMissingPositive([]int{3, 4, -1, 1})
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case3() {
	ans := firstMissingPositive([]int{7, 8, 9, 11, 12})
	fmt.Print(ans)
	// Output:
	// 1
}
