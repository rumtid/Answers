package leetcode

import "fmt"

func Example_case1() {
	ans := swapPairs(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
	)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}
	// Output:
	// 2->1->4->3
}
