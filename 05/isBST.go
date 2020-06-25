	//判断一个二叉树是否为搜索二叉树
func isValidBST(root *TreeNode) bool{
	return dfs(root, -1 << 63, 1 << (63-1))
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max{
		return false
	}
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}