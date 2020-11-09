package leetcode

import "fmt"

func Example_case1() {
	ans := simplifyPath("/home/")
	fmt.Print(ans)
	// Output:
	// /home
}

func Example_case2() {
	ans := simplifyPath("/../")
	fmt.Print(ans)
	// Output:
	// /
}

func Example_case3() {
	ans := simplifyPath("/home//foo/")
	fmt.Print(ans)
	// Output:
	// /home/foo
}

func Example_case4() {
	ans := simplifyPath("/a/./b/../../c/")
	fmt.Print(ans)
	// Output:
	// /c
}

func Example_case5() {
	ans := simplifyPath("/a/../../b/../c//.//")
	fmt.Print(ans)
	// Output:
	// /c
}

func Example_case6() {
	ans := simplifyPath("/a//b////c/d//././/..")
	fmt.Print(ans)
	// Output:
	// /a/b/c
}
