package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	p := &head
	for i := 0; i < m-1; i++ {
		p = &(*p).Next
	}
	l1, l2 := *p, (*p).Next
	for i := 0; i < n-m; i++ {
		l1, l2, l2.Next = l2, l2.Next, l1
	}
	*p, (*p).Next = l1, l2
	return head
}
