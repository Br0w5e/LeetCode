package main

//83. 删除排序链表中的重复元素

//82. 删除排序链表中的重复元素 II

//删除有序链表中的重复元素(保留一个)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre, cur := head, head.Next
	for cur != nil {
		//相同值，忽略cur节点
		if pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			// pre = pre.Next
			//向后移动到cur
			pre = cur
		}
		//cur向后移动一个节点
		cur = cur.Next
	}

	return head
}

//删除有序链表中的重复元素(不保留)
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyNode.Next
}
