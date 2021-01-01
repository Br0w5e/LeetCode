package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isValidBST(root *TreeNode) bool {
	return dfs(root, (-1)<<63, 1<<(63-1))
}

func dfs(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}
