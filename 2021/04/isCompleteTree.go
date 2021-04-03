package main

//958. 二叉树的完全性检验

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//层次遍历，记得标记
func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			flag = true
			continue
		}

		if flag {
			return false
		}

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return true
}
