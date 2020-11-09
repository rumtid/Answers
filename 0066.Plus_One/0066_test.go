package leetcode

import "fmt"

func Example_case1() {
	ans := plusOne([]int{1, 2, 3})
	fmt.Print(ans)
	// Output:
	// [1 2 4]
}

func Example_case2() {
	ans := plusOne([]int{4, 3, 2, 1})
	fmt.Print(ans)
	// Output:
	// [4 3 2 2]
}
