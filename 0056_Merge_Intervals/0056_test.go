package main

import "fmt"

func Example_case1() {
	ans := merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	})
	fmt.Print(ans)
	// Output:
	// [[1 6] [8 10] [15 18]]
}

func Example_case2() {
	ans := merge([][]int{
		[]int{1, 4},
		[]int{4, 5},
	})
	fmt.Print(ans)
	// Output:
	// [[1 5]]
}
