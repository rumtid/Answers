package leetcode

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for i := 0; i < len(ans); i++ {
		sort.Ints(ans[i])
		fmt.Println(ans[i])
	}
	// Unordered output:
	// [-1 0 1]
	// [-1 -1 2]
}
