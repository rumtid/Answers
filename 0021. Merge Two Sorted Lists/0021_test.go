package main

import "fmt"

func Example_case1() {
	ans := mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
	)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}
	// Output:
	// 1->1->2->3->4->4
}
