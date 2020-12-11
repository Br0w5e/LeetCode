package main

//1008. 前序遍历构造二叉搜索树
import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	inorder := make([]int, len(preorder))
	copy(inorder, preorder)
	sort.Ints(inorder)
	return buildTree(preorder, inorder)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	res := &TreeNode {
		Val : preorder[0],
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

func bstFromPreorder1(preorder []int) *TreeNode  {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	i := 1
	for ; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			break
		}
	}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}