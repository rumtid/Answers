package leetcode

import "fmt"

func Example_case1() {
	ans := generateMatrix(3)
	for _, row := range ans {
		fmt.Println(row)
	}
	// Output:
	// [1 2 3]
	// [8 9 4]
	// [7 6 5]
}
