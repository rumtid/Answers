package main

import "fmt"

func Example_case1() {
	ans := minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	})
	fmt.Print(ans)
	// Output:
	// 7
}
