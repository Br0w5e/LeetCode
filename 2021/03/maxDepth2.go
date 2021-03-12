package main

//559. N 叉树的最大深度

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//bfs
func maxDepth(root *Node) int {
	res := 0
	if root == nil {
		return res
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			queue = append(queue, node.Children...)
			// for _, child := range node.Children {
			//     if child != nil {
			//         queue = append(queue, child)
			//     }
			// }
		}
	}
	return res
}

//dfs
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil {
		return 1
	}
	maxdepth := 0
	for _, child := range root.Children {
		maxdepth = max(maxdepth, maxDepth(child))
	}
	return maxdepth+1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}