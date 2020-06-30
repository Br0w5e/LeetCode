package main

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val+leftGain+rightGain

		maxSum = max(maxSum, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	doMaxPathSum(root, &result)
	return result
}

func doMaxPathSum(node *TreeNode, max *int) int {
	if node == nil {
		return math.MinInt32
	}
	leftMax := doMaxPathSum(node.Left, max)
	rightMax := doMaxPathSum(node.Right, max)
	if leftMax < 0 {
		leftMax = 0
	}
	if rightMax < 0 {
		rightMax = 0
	}
	var result int
	if rightMax > leftMax {
		result = node.Val+rightMax
	} else {
		result = node.Val+leftMax
	}
	all := node.Val+leftMax+rightMax
	if all > *max {
		*max = all
	}
	return result
}
