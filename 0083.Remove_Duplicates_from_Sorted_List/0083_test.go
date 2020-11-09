package leetcode

import "fmt"

func Example_case1() {
	ans := deleteDuplicates(
		&ListNode{1, &ListNode{1, &ListNode{2, nil}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->2
}

func Example_case2() {
	ans := deleteDuplicates(
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->2->3
}
