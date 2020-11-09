package leetcode

import "fmt"

func Example_case1() {
	ans := removeNthFromEnd(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		2,
	)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}
	// Output:
	// 1->2->3->5
}
