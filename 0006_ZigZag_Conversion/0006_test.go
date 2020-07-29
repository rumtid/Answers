package main

import "fmt"

func Example_case1() {
	ans := convert("PAYPALISHIRING", 3)
	fmt.Print(ans)
	// Output:
	// PAHNAPLSIIGYIR
}

func Example_case2() {
	ans := convert("PAYPALISHIRING", 4)
	fmt.Print(ans)
	// Output:
	// PINALSIGYAHRPI
}
