package main

//872. 叶子相似的树
//dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := make([]int, 0)
	leaf2 := make([]int, 0)
	getLeaf(root1, &leaf1)
	getLeaf(root2, &leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func getLeaf(root *TreeNode, nums *[]int) {
	if root.Left == nil && root.Right == nil {
		*nums = append(*nums, root.Val)
		return
	}
	if root.Left != nil {
		getLeaf(root.Left, nums)
	}
	if root.Right != nil {
		getLeaf(root.Right, nums)
	}
}
