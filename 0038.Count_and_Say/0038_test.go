package leetcode

import "fmt"

func Example_case1() {
	ans := countAndSay(1)
	fmt.Print(ans)
	// Output:
	// 1
}

func Example_case2() {
	ans := countAndSay(4)
	fmt.Print(ans)
	// Output:
	// 1211
}
