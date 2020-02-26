package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var p1, p2, p3 **ListNode = &head, nil, nil
	for {
		if *p1 == nil || (*p1).Next == nil {
			break
		}
		p2, p3 = &(*p1).Next, &(*p1).Next.Next
		*p1, *p2, *p3 = *p2, *p3, *p1
		p1 = p2
	}
	return head
}
