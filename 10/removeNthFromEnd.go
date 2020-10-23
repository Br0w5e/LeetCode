package main

//19. 删除链表的倒数第N个节点
//不设置哑结点的方式
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	//节点数等于n，相当于删除了头结点，那么应该返回头结点的下一个节点
	if fast == nil {
		head = head.Next
		return head
	}
	//需要将slow指向倒数第n+1个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//删除倒数第n个节点
	slow.Next = slow.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for n > 0 || fast != nil {
		if n > 0 {
			fast = fast.Next
			n--
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}