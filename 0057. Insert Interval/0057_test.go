package main

import "fmt"

func Example_case1() {
	ans := insert([][]int{
		[]int{1, 3},
		[]int{6, 9},
	}, []int{2, 5})
	fmt.Print(ans)
	// Output:
	// [[1 5] [6 9]]
}

func Example_case2() {
	ans := insert([][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}, []int{4, 8})
	fmt.Print(ans)
	// Output:
	// [[1 2] [3 10] [12 16]]
}
