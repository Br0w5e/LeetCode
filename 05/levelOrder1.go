package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		res = append(res, make([]int, 0))
		n := len(queue)
		for i := 0; i < n; i++ {
			top := queue[0]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			res[ans] = append(res[ans], top.Val)
			queue = queue[1:]
		}
		ans++
	}
	return res
}
