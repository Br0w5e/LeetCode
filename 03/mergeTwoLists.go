// 合并两个有序链表为一个
func merge(A []int, m int, B []int, n int) []int {
    i, j, k := m-1, n-1, m+n-1
    for j >= 0 {
        if i >= 0 && A[i] > B[j] {
            A[k] = A[i]
            i--
        } else {
            A[k] = B[j]
            j--
        }
        k--
    }
    return A
}
// 进行递归排序，  别
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    }
    l2.Next = mergeTwoLists(l1, l2.Next)
    return l2
}

//进行迭代排序
//别忘了浮动单节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, node *ListNode
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        head = l1
        node = l1
        l1 = l1.Next
    } else {
        head = l2
        node = l2
        l2 = l2.Next
    }
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            node.Next = l1
            l1 = l1.Next
        } else {
            node.Next = l2
            l2 = l2.Next
        }
        node = node.Next
    }
    if l1 != nil {
        node.Next = l1
    } else {
        node.Next = l2
    }
    return head
}