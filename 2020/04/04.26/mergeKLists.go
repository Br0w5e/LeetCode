//合并k个有序链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    l := len(lists)
    if l == 0 {
        return nil
    }
    if l == 1 {
        return lists[0]
    }
    
    mid := l/2
    left := mergeKLists(lists[:mid])
    right := mergeKLists(lists[mid:])
    return mergeSort(left, right)
}

func mergeSort(left, right *ListNode) *ListNode {
    // 两个链表有一个为空时
    if left == nil || right == nil {
        if left == nil {
            return right
        }
        if right == nil {
            return left
        }
        return nil
    }
    retPre := &ListNode{}
    cur := retPre
    for left != nil && right != nil {
        if left.Val < right.Val {
            cur.Next = left
            left = left.Next
        } else {
            cur.Next = right
            right = right.Next
        }
        cur = cur.Next // 链表当前指针后移
    }
    if left != nil {
        cur.Next = left
    }
    if right != nil {
        cur.Next = right
    }
    return retPre.Next
}