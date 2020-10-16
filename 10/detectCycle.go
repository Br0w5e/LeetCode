//检测链表是否有环，有环的时候返回该节点
//142. 环形链表 II
//参考hasCycle.go

//双指针
package main
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

//哈希表
//当节点在哈希表中第二次出现的返回节点，否则标记为true
func detectCycle2(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if m[node] {
			return node
		}
		m[node] = true
		node = node.Next
	}
	return nil
}