package main

import "sort"

//148. 排序链表
//数组方法
func sortList(head *ListNode) *ListNode {
	values := make([]int, 0)
	p := head
	for p != nil {
		values = append(values, p.Val)
		p = p.Next
	}
	p = head
	sort.Ints(values)
	for p != nil {
		p.Val = values[0]
		values = values[1:]
		p = p.Next
	}
	return head
}

//归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	//左右两边分别进行归并
	lnode, rnode := sortList(head), sortList(mid)
	//对链表进行合并
	return merge2SortedListNode(lnode, rnode)
}

//合并链表
func merge2SortedListNode(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = merge2SortedListNode(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge2SortedListNode(l1, l2.Next)
		return l2
	}
}
