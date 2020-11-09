package leetcode

import "fmt"

func Example_case1() {
	ans := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Print(ans)
	// Output:
	// 6
}
