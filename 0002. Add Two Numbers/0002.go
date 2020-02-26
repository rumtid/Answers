package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	p, v := &l3, 0
	for l1 != nil || l2 != nil || v != 0 {
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		*p = &ListNode{v % 10, nil}
		p, v = &((*p).Next), v/10
	}
	return l3
}
