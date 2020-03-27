package main

import "fmt"

func Example_case1() {
	ans := maximalRectangle([][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	})
	fmt.Print(ans)
	// Output:
	// 6
}
