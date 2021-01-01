//移除链表中的重复节点
//map标记

package  main

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := head
	m := make(map[int]bool)
	m[head.Val] = true
	for head.Next != nil {
		_, ok := m[head.Next.Val]
		//已经存在
		if ok {
			head.Next = head.Next.Next
		} else {
			m[head.Next.Val] = true
			head = head.Next
		}
	}
	return res
}
