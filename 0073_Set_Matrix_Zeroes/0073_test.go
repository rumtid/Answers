package main

import "fmt"

func Example_case1() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
	// Output:
	// [1 0 1]
	// [0 0 0]
	// [1 0 1]
}

func Example_case2() {
	matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
	// Output:
	// [0 0 0 0]
	// [0 4 5 0]
	// [0 3 1 0]
}
