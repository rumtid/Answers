package leetcode

import "fmt"

func Example_case1() {
	ans := reverseList(
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
	)
	for ; ans != nil; ans = ans.Next {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
	}
	// Output:
	// 5->4->3->2->1
}
