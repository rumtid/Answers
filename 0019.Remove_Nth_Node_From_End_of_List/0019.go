package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	for i := 0; i < n; i++ {
		tail = tail.Next
	}

	pp := &head
	for tail != nil {
		tail = tail.Next
		pp = &(*pp).Next
	}

	*pp = (*pp).Next

	return head
}
