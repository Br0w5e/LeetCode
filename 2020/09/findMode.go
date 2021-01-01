package main

//501. 二叉搜索树中的众数

//要求不开辟空间，这里开辟空间了

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	dfs(root, m)
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	res := make([]int, 0)
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func dfs(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	dfs(root.Left, m)
	m[root.Val]++
	dfs(root.Right, m)
}
