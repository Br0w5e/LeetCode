package main

//513. 找树左下角的值


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//层次遍历，内存占用大
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return res[len(res)-1][0]
}

//在层次遍历的时候先考虑右边节点，再考虑左边节点，则队列中最后一个节点就是最左边的节点
//稍作改变，在层次遍历的时候先考虑左边节点，再考虑右边节点，则在队列中最后一个节点就是最后一层最右边节点
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := -1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = node.Val
		// 注意，一定要先遍历右节点，再遍历左节点
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

	}
	return res
}
