//检测链表是否有环，有环的时候返回该节点
//参考hasCycle.go
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil &&	fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head 
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}