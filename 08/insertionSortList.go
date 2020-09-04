package main

import "sort"

type ListNode struct {
	Val int
	Next *ListNode
}
func insertionSortList(head *ListNode) *ListNode  {
	if head == nil {
		return nil
	}

	h := &ListNode{}
	p := head.Next
	h.Next = head
	head.Next = nil
	for p != nil {
		prev := h
		next := p.Next

		//寻找插入点
		for prev.Next != nil && prev.Next.Val <= p.Val {
			prev = prev.Next
		}
		p.Next = prev.Next
		prev.Next = p
		p = next
	}
	return h.Next
}
//野路子
func insertionSortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nums := make([]int,0)
	for head != nil {
		nums = append(nums,head.Val)
		head = head.Next
	}
	sort.Ints(nums)
	rst := &ListNode{}
	tmp := rst
	for _,v := range nums {
		tmp.Next = &ListNode{Val:v}
		tmp = tmp.Next
	}
	return rst.Next
}