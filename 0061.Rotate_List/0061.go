package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var n int = 1
	var p *ListNode
	for p = head; p.Next != nil; p = p.Next {
		n++
	}
	k = n - k%n
	p.Next = head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	head = p.Next
	p.Next = nil
	return head
}
