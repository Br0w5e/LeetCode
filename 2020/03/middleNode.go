//返回List的中间节点
//思路1：先计数再返回
func middleNode(head *ListNode) *ListNode {
	p := head
	cnt := 0
	for p != nil {
		p = p.Next
		cnt++
	}
	for i := 0; i < cnt/2; i++{
		head = head.Next
	}
	return head
}

//思路二：双指针
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head 
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}