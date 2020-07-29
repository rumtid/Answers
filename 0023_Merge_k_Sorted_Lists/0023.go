package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	l := (*h)[len(*h)-1]
	(*h)[len(*h)-1] = nil
	*h = (*h)[:len(*h)-1]
	return l
}

func mergeKLists(lists []*ListNode) *ListNode {
	for i := 0; i < len(lists); {
		if lists[i] == nil {
			lists[i] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		} else {
			i++
		}
	}

	var listheap ListHeap = lists
	heap.Init(&listheap)
	var head *ListNode
	p := &head

	for listheap.Len() != 0 {
		node := heap.Pop(&listheap).(*ListNode)
		*p = node
		p = &(*p).Next
		if node.Next != nil {
			heap.Push(&listheap, node.Next)
		}
	}

	return head
}
