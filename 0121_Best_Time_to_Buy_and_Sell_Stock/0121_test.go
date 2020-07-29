package main

import (
	"fmt"
)

func Example_case1() {
	ans := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Print(ans)
	// Output:
	// 5
}

func Example_case2() {
	ans := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Print(ans)
	// Output:
	// 0
}
