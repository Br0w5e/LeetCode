package main




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//515. 在每个树行中找最大值

//层次遍历，每层记录最大值即可
//bfs
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	//用来存储每一层的节点和下一层节点
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		max := node.Val
		n := len(queue)
		//层次遍历
		for i := 0; i < n; i++ {
			node = queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Val > max {
				max = node.Val
			}
		}
		//同一层的出队
		queue = queue[n:]
		res = append(res, max)
	}
	return res
}