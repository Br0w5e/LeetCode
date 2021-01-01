/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//206. 反转链表
package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next //暂存cur.Next
		cur.Next = pre   //将cur.Next变为pre
		pre = cur        //pre向后移动一位
		cur = temp       //将cur置位temp
	}
	return pre
}

