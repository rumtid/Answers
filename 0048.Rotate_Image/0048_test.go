package leetcode

import "fmt"

func Example_case1() {
	img := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	rotate(img)
	for _, row := range img {
		fmt.Println(row)
	}
	// Output:
	// [7 4 1]
	// [8 5 2]
	// [9 6 3]
}

func Example_case2() {
	img := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(img)
	for _, row := range img {
		fmt.Println(row)
	}
	// Output:
	// [15 13 2 5]
	// [14 3 4 1]
	// [12 6 8 9]
	// [16 7 10 11]
}
