package main

//563. 二叉树的坡度


func findTilt(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	*res += abs(right-left)
	return left + right + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
