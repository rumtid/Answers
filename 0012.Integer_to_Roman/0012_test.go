package leetcode

import "fmt"

func Example_case1() {
	ans := intToRoman(3)
	fmt.Print(ans)
	// Output:
	// III
}

func Example_case2() {
	ans := intToRoman(4)
	fmt.Print(ans)
	// Output:
	// IV
}

func Example_case3() {
	ans := intToRoman(9)
	fmt.Print(ans)
	// Output:
	// IX
}

func Example_case4() {
	ans := intToRoman(58)
	fmt.Print(ans)
	// Output:
	// LVIII
}

func Example_case5() {
	ans := intToRoman(1994)
	fmt.Print(ans)
	// Output:
	// MCMXCIV
}
