package leetcode

import "fmt"

func Example_case1() {
	ans := myPow(2.0, 10)
	fmt.Printf("%.6g", ans)
	// Output:
	// 1024
}

func Example_case2() {
	ans := myPow(2.1, 3)
	fmt.Printf("%.6g", ans)
	// Output:
	// 9.261
}

func Example_case3() {
	ans := myPow(2.0, -2)
	fmt.Printf("%.6g", ans)
	// Output:
	// 0.25
}
