package main

import "fmt"

func Example_case1() {
	ans := romanToInt("III")
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	ans := romanToInt("IV")
	fmt.Print(ans)
	// Output:
	// 4
}

func Example_case3() {
	ans := romanToInt("IX")
	fmt.Print(ans)
	// Output:
	// 9
}

func Example_case4() {
	ans := romanToInt("LVIII")
	fmt.Print(ans)
	// Output:
	// 58
}

func Example_case5() {
	ans := romanToInt("MCMXCIV")
	fmt.Print(ans)
	// Output:
	// 1994
}
