package leetcode

import "fmt"

func Example_case1() {
	ans := partition(
		&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}},
		3,
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->2->2->4->3->5
}
