package leetcode

import (
	"fmt"

	"github.com/rumtid/leetcode/utils"
)

func Example_case1() {
	ans := generateTrees(3)
	for _, root := range ans {
		fmt.Println(utils.Ntoa(root))
	}
	// Unordered output:
	// [1 [] [3 [2] []]]
	// [3 [2 [1] []] []]
	// [3 [1 [] [2]] []]
	// [2 [1] [3]]
	// [1 [] [2 [] [3]]]
}
