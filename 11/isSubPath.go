package main

//1367. 二叉树中的列表
//双层深度遍历
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if dfs(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
func dfs(head *ListNode, root *TreeNode) bool {
	//链表遍历完毕
	if head == nil {
		return true
	}
	//链表未遍历完毕，但树已经完毕
	if root == nil {
		return false
	}
	//链表的值和树的值不相等
	if root.Val != head.Val {
		return false
	}
	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
}
