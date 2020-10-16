package main

//交换链表中的相邻节点
//24. 两两交换链表中的节点
//递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

//设置哑结点
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}
	tmp := dummy
	for tmp.Next != nil && tmp.Next.Next != nil {
		//tmp -> first -> second
		first, second := tmp.Next, tmp.Next.Next
		tmp.Next = second
		first.Next = second.Next
		second.Next = first
		//上一步交换之后的节点关系已经变成 tmp -> second -> first
		tmp = first
	}
	return dummy.Next
}