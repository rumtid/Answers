package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[3 [9] [20 [15] [7]]]"))
	ans := maxDepth(root)
	fmt.Print(ans)
	// Output:
	// 3
}

func Example_case2() {
	root := (*TreeNode)(utils.Aton("[1 [] [2]]"))
	ans := maxDepth(root)
	fmt.Print(ans)
	// Output:
	// 2
}

func Example_case3() {
	root := (*TreeNode)(utils.Aton("[]"))
	ans := maxDepth(root)
	fmt.Print(ans)
	// Output:
	// 0
}

func Example_case4() {
	root := (*TreeNode)(utils.Aton("[0]"))
	ans := maxDepth(root)
	fmt.Print(ans)
	// Output:
	// 1
}
