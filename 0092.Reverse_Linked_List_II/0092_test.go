package leetcode

import "fmt"

func Example_case1() {
	ans := reverseBetween(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		2, 4,
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->4->3->2->5
}
