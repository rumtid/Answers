package leetcode

import "fmt"

func Example_case1() {
	ans := searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}, 3)
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}, 13)
	fmt.Print(ans)
	// Output:
	// false
}
