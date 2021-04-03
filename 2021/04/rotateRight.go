package main

//61. 旋转链表

//将链表向右旋转n次
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//计算链表长度
	cnt := 1
	c := head
	for c.Next != nil {
		c = c.Next
		cnt++
	}
	//旋转操作
	for i := 0; i < k%cnt; i++ {
		next := head.Next
		cur := head
		for next.Next != nil {
			next = next.Next
			cur = cur.Next
		}
		cur.Next = nil
		next.Next = head
		head = next
	}
	return head
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	//计数
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	//移动次数
	add := n - k%n
	if add == n {
		return head
	}
	//移动k次，iter当前指向尾结点
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	//断开地方
	ret := iter.Next
	iter.Next = nil
	return ret
}
