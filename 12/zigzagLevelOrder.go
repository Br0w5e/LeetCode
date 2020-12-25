package main

//103. 二叉树的锯齿形层序遍历
//锯齿形遍历 参考levelOrder.go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	level := 0
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		res = append(res, make([]int, 0))
		n := len(queue)
		for i := 0; i < n; i++ {
			top := queue[i]
			res[level] = append(res[level], top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		queue = queue[n:]
		if level%2 == 1 {
			reverse(res[level])
		}
		level++
	}
	return res
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
