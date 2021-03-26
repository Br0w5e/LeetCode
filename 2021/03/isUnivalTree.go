package main

//965. 单值二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//bfs
func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Val != val {
			return false
		}
	}
	return true
}

//dfs
func isUnivalTree(root *TreeNode) bool {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return dfs(node.Left, val) && dfs(node.Right, val)
}
