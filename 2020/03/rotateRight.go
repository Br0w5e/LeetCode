//将链表向右旋转n次
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    //计算链表长度
    cnt := 1
    c := head 
    for c.Next != nil {
        c = c.Next
        cnt++
    }
    //旋转操作
    for i := 0; i < k%cnt; i++ {
        next := head.Next
        cur := head
        for next.Next != nil {
            next = next.Next
            cur = cur.Next
        }
        cur.Next = nil
        next.Next = head
        head = next
    }
    return head
}