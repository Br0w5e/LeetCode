package main

//1325. 删除给定值的叶子节点
//递归，后序遍历吧
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	//递归出口
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	//叶子节点，并且值相等, 断开该节点
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
