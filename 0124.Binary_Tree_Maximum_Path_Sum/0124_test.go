package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[1 [2] [3]]"))
	ans := maxPathSum(root)
	fmt.Print(ans)
	// Output:
	// 6
}

func Example_case2() {
	root := (*TreeNode)(utils.Aton("[-10 [9] [20 [15] [7]]]"))
	ans := maxPathSum(root)
	fmt.Print(ans)
	// Output:
	// 42
}
