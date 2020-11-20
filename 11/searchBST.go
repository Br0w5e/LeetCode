package main

//700. 二叉搜索树中的搜索
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//树的层次遍历
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		if node.Val == val {
			return node
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return nil
}

//迭代，结合二叉树的性质
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}
