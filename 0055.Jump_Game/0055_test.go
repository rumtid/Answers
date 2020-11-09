package leetcode

import "fmt"

func Example_case1() {
	ans := canJump([]int{2, 3, 1, 1, 4})
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := canJump([]int{3, 2, 1, 0, 4})
	fmt.Print(ans)
	// Output:
	// false
}
