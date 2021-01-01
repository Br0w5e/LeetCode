package main
//面试题 04.05. 合法二叉搜索树

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

//性质
func isValidBST2(root *TreeNode) bool {
	res := make([]int, 0)
	dfs(root, &res)
	for i := 0; i < len(res)-1; i++ {
		if res[i+1] <= res[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}