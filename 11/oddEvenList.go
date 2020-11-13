package main

//328. 奇偶链表
//数组方法
func oddEvenList(head *ListNode) *ListNode {
	values := make([]int, 0)
	p := head
	for p != nil {
		values = append(values, p.Val)
		p = p.Next
	}

	p = head
	for i := 0; i < len(values); i++ {
		if i % 2 == 0 {
			p.Val = values[i]
			p = p.Next
		}
	}
	for i := 1; i < len(values); i++ {
		if i % 2 == 1 {
			p.Val = values[i]
			p = p.Next
		}
	}
	return head
}

//链表方法
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//从第三个节点开始直到链表末尾
	//奇偶活动节点
	odd := head
	even := head.Next
	//偶数节点的头结点
	startEven := even
	//当前遍历节点
	current := even.Next
	//记录奇偶变化
	flag := false
	for current != nil {
		if !flag {
			//奇数节点（从1开始）
			odd.Next = current
			odd = odd.Next
		} else {
			//奇数节点
			even.Next = current
			even = even.Next
		}
		//下移节点，并变化标志
		current = current.Next
		flag = !flag
	}
	if flag {
		//节点的数量是奇数个，将偶数链的最后一个节点指向空节点（否则该节点将指向倒数第一个节点）
		even.Next = nil
	}
	odd.Next = startEven
	return head
}
