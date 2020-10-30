package main

//86. 分隔链表
//用两个数组分别存储
func partition(head *ListNode, x int) *ListNode {
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	p := head
	for p != nil {
		if p.Val < x {
			nums1 = append(nums1, p.Val)
		} else {
			nums2 = append(nums2, p.Val)
		}
		p = p.Next
	}
	nums1 = append(nums1, nums2...)
	p = head
	for i := 0; i < len(nums1); i++ {
		p.Val = nums1[i]
		p = p.Next
	}
	return head
}
//两个链表，最后链接
func partition(head *ListNode, x int) *ListNode {
	//输入: head = 1->4->3->2->5->2, x = 3
	//输出: 1->2->2->4->3->5
	if head == nil{
		return head
	}
	headDumy := &ListNode{Val: 0}
	tailDumy := &ListNode{Val: 0}
	headDumy.Next = head
	head = headDumy
	tail := tailDumy
	head = headDumy
	for head.Next != nil{
		if head.Next.Val < x{
			head = head.Next
		}else{
			//移除当前节点
			tmp := head.Next
			head.Next = head.Next.Next
			tail.Next = tmp
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDumy.Next
	return headDumy.Next

}