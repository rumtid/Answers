package main

import "fmt"

func Example_case1() {
	ans := uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	})
	fmt.Print(ans)
	// Output:
	// 2
}
