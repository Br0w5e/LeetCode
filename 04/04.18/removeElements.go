//删除链表中的指定节点
//设置头节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{
        Val : -1,
        Next : head,
    }
    node := preHead
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		}else {
			node = node.Next
		}
	}
	return preHead.Next
}
