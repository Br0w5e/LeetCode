//递归，了解二叉树的性质

package main

// 回复二叉树， 由前序和中序回复
// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	res := &TreeNode{
		Val: preorder[0],
	}
	if n == 1 {
		return res
	}
	index := getIndex(inorder, res.Val)
	res.Left = buildTree(preorder[1:index+1], inorder[:index])
	res.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return res
}

func getIndex(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return -1
}

//106. 从中序与后序遍历序列构造二叉树
//由中序和后序回复
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	res := &TreeNode{
		Val: postorder[n-1],
	}
	if n == 1 {
		return res
	}
	index := getIndex(inorder, res.Val)
	res.Left = buildTree(inorder[:index], postorder[:index])
	res.Right = buildTree(inorder[index+1:], postorder[index:n-1])
	return res
}

func getIndex(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return 0
}
