package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var l1 *ListNode
	for l2 := head; l2 != nil; {
		l1, l2, l2.Next = l2, l2.Next, l1
	}
	return l1
}
