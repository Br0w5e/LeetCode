package main

//538. 把二叉搜索树转换为累加树
// 反序中序遍历
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	dfs(root, &res)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root != nil {
		//遍历右子树
		dfs(root.Right, sum)
		//遍历顶节点
		root.Val += *sum
		*sum = root.Val
		//遍历左子树
		dfs(root.Left, sum)
		return
	}
	return
}

//1038. 从二叉搜索树到更大和树
func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root != nil {
		dfs(root.Right, sum)
		root.Val += *sum
		*sum = root.Val
		dfs(root.Left, sum)
		return
	}
	return
}
