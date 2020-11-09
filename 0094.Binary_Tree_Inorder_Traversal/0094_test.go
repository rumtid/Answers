package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[1 [] [2 [3] []]]"))
	ans := inorderTraversal(root)
	fmt.Print(ans)
	// Output:
	// [1 3 2]
}
