package leetcode

import "fmt"

func Example_case1() {
	ans := grayCode(2)
	fmt.Print(ans)
	// Output:
	// [0 1 3 2]
}

func Example_case2() {
	ans := grayCode(0)
	fmt.Print(ans)
	// Output:
	// [0]
}
