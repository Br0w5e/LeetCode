package main

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
