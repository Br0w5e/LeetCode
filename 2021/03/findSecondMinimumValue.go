package main

//671. 二叉树中第二小的节点

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//dfs
//问题转化为求左右子树的最小值，如果左右子树最小值都大于根节点的值取较小的值。其他情况取左右子树较大的值。
func findSecondMinimumValue(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val > val {
		return root.Val
	}
	l := dfs(root.Left, val)
	r := dfs(root.Right, val)
	if l > val && r > val {
		return min(l, r)
	}
	return max(l, r)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
