package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[3 [9] [20 [15] [7]]]"))
	ans := zigzagLevelOrder(root)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// Output:
	// [3]
	// [20 9]
	// [15 7]
}
