//移除链表中的重复节点
//map标记
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := head 
	listMap := make(map[int]bool)
	listMap[head.Val] = true
	for head.Next != nil {
		_, ok := listMap[head.Next.Val]
		//已经标记
		if ok {
			head.Next = head.Next.Next
		} else {
			listMap[head.Next.Val] = true
			head = head.Next
		}
	}
	return res
}