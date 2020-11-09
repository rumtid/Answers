package leetcode

import "fmt"

func Example_case1() {
	ans := generateParenthesis(3)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// Output:
	// ((()))
	// (()())
	// (())()
	// ()(())
	// ()()()
}
