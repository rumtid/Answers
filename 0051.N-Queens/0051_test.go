package leetcode

import "fmt"

func Example_case1() {
	ans := solveNQueens(4)
	for _, row := range ans {
		fmt.Println(row)
	}
	// Unordered output:
	// [.Q.. ...Q Q... ..Q.]
	// [..Q. Q... ...Q .Q..]
}
