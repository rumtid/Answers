package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}
