package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	p3 := &l3
	for {
		if l1 == nil {
			*p3 = l2
			break
		}
		if l2 == nil {
			*p3 = l1
			break
		}
		if l1.Val > l2.Val {
			*p3 = l2
			l2 = l2.Next
			p3 = &(*p3).Next
		} else {
			*p3 = l1
			l1 = l1.Next
			p3 = &(*p3).Next
		}
	}
	return l3
}
