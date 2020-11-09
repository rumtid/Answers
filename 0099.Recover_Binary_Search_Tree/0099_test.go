package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[1 [3 [] [2]] []]"))
	recoverTree(root)
	fmt.Print(utils.Ntoa(root))
	// Output:
	// [3 [1 [] [2]] []]
}

func Example_case2() {
	root := (*TreeNode)(utils.Aton("[3 [1] [4 [2] []]]"))
	recoverTree(root)
	fmt.Print(utils.Ntoa(root))
	// Output:
	// [2 [1] [4 [3] []]]
}
