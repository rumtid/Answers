package main

import "fmt"

func Example_case1() {
	ans := reverseKGroup(
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
	// 2->1->4->3->5
}

func Example_case2() {
	ans := reverseKGroup(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		3,
	)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}
	// Output:
	// 3->2->1->4->5
}
