package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var l1, l2 *ListNode
	p1, p2 := &l1, &l2
	for ; head != nil; head = head.Next {
		if head.Val < x {
			*p1 = head
			p1 = &head.Next
		} else {
			*p2 = head
			p2 = &head.Next
		}
	}
	*p1, *p2 = l2, nil
	return l1
}
