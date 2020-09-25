package main

// 617. 合并二叉树
//规则同一位置节点进行相加

//方法：递归
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := &TreeNode{
		Val: t1.Val + t2.Val,
	}
	root.Left = mergeTrees2(t1.Left, t2.Left)
	root.Right = mergeTrees2(t1.Right, t2.Right)
	return root
}
