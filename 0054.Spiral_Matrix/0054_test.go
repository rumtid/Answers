package leetcode

import "fmt"

func Example_case1() {
	ans := spiralOrder([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	})
	fmt.Print(ans)
	// Output:
	// [1 2 3 6 9 8 7 4 5]
}

func Example_case2() {
	ans := spiralOrder([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	})
	fmt.Print(ans)
	// Output:
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}
