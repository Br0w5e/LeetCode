package main
//判断链表是否有环

//判断链表是否有环
//快慢指針 推荐
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
        if fast == slow {
            return true
        }
	}
	return false
}

//判断链表是否有环，并返回环的入口
//面试题 02.08. 环路检测
//快慢指针，有性质
func detectCycle1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    slow := head
    fast := head
    for fast.Next != nil && fast.Next.Next != nil {
        //走完了
        fast = fast.Next.Next
        slow = slow.Next
        //相遇
        if fast == slow {
            //head 开始走
            for head != fast {
                head = head.Next
                fast = fast.Next
            }
            return head
        }
    }
    return nil
}