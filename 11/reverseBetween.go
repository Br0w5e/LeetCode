package main

//92. 反转链表 II
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	//要翻转的头结点
	head = pre.Next
	//翻转
	for i := m; i < n ; i++ {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return dummy.Next
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	if head == nil {
		return nil
	}

	// 头部变化所以使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	// 最开始：0->1->2->3->4->5->nil
	var pre *ListNode
	for i := 0; i < m; i++ {
		pre = head
		head = head.Next
	}
	// 遍历之后： 1(pre)->2(head)->3->4->5->NULL
	var next *ListNode
	// 用于中间节点连接
	var mid = head
	for j := 0; head != nil && j <= n-m; j++ {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		// 循环中 1(pre).Next 一直是 2(mid)
		// 第一次循环： 1(pre) nil<-2(mid,next) 3(tmp,head)->4->5->nil
		// 第一次循环： 1(pre) nil<-2(mid)<-3(next) 4(tmp,head)->5->nil

	}
	// 循环需要执行四次
	// 循环结束：1(pre) nil<-2(mid)<-3<-4(next) 5(head)->nil
	pre.Next = next
	mid.Next = head
	return dummy.Next
}
