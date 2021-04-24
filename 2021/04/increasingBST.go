package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	dfs(root, &nums)
	res := &TreeNode{}
	node := res
	for len(nums) > 0 {
		node.Val = nums[0]
		if len(nums) == 1 {
			break
		}
		node.Right = &TreeNode{}
		node = node.Right
		nums = nums[1:]
	}
	return res
}

func dfs(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, nums)
	*nums = append(*nums, root.Val)
	dfs(root.Right, nums)
}
