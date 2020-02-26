package main

import "fmt"

func Example_case1() {
	ans := addTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
	}
	// Output:
	// 708
}
