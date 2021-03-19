package main
//面试题 02.07. 链表相交

//方法 a + mid + b = b + mid + a  最后共同走过的路才是一致的
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	na, nb := headA, headB

	for na != nb {
		if na != nil {
			na = na.Next
		} else {
			na = headB
		}

		if nb != nil {
			nb = nb.Next
		} else {
			nb = headA
		}
	}

	return na;
}

