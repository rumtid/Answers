package leetcode

import "fmt"

func Example_case1() {
	ans := rotateRight(
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
	// 4->5->1->2->3
}

func Example_case2() {
	ans := rotateRight(
		&ListNode{0, &ListNode{1, &ListNode{2, nil}}},
		4,
	)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}
	// Output:
	// 2->0->1
}
