//返回链中倒数第K个节点
//思路：快慢指针
func kthToLast(head *ListNode, k int) int {
    fast := head
    slow := head
    for ; k != 0; k-- {
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    return slow.Val
}