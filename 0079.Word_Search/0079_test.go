package leetcode

import "fmt"

var board = [][]byte{
	[]byte{'A', 'B', 'C', 'E'},
	[]byte{'S', 'F', 'C', 'S'},
	[]byte{'A', 'D', 'E', 'E'},
}

func Example_case1() {
	ans := exist(board, "ABCCED")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case2() {
	ans := exist(board, "SEE")
	fmt.Print(ans)
	// Output:
	// true
}

func Example_case3() {
	ans := exist(board, "ABCB")
	fmt.Print(ans)
	// Output:
	// false
}
