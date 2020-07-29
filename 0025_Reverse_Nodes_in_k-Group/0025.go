package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	prev := head
	next := head.Next
	for next != tail {
		prev, next, next.Next = next, next.Next, prev
	}
	head.Next = reverseKGroup(tail, k)
	return prev
}
