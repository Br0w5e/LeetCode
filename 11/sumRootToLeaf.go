package main

//1022. 从根到叶的二进制数之和
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	//出口
	if root == nil {
		return 0
	}
	//求和
	sum = sum*2 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	//回溯
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}