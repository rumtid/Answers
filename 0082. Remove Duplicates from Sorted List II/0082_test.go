package main

import "fmt"

func Example_case1() {
	ans := deleteDuplicates(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->2->5
}

func Example_case2() {
	ans := deleteDuplicates(
		&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 2->3
}
