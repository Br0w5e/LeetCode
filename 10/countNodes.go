package main

//222. 完全二叉树的节点个数

func countNodes(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	*res++
	dfs(root.Left, res)
	dfs(root.Right, res)
}