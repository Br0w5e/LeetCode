//删除有序链表中的重复元素(保留一个)
func  deleteDuplicates(head *TreeNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head 
	cur := head.Next 
	for cur != nil {
		if pre.Val == cur.Val{
			pre.Next == cur.Next
		} else {
			//也可以是 pre = pre.Next
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
//删除有序链表中的重复元素(不保留)
func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	cur := pre
	count := 0
	for head != nil && head.Next != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			count++
			head = head.Next
		}
		if count > 0 {
			cur.Next = head.Next
			head = head.Next
			count = 0
		} else {
			cur = head
			head = head.Next
		}
	}
	return pre.Next
}