//返回树的最大深度
// 方法：递归
package main
//DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Right), maxDepth(root.Left))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Val   *TreeNode
	depth int
}
//BFS
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num := 0
	queue := []*Node{}
	queue = append(queue, &Node{root, 1})
	for len(queue) != 0 {
		root1 := queue[0]
		queue = queue[1:]
		num = max(num, root1.depth)
		if root1.Val.Left != nil {
			queue = append(queue, &Node{root1.Val.Left, root1.depth + 1})
		}
		if root1.Val.Right != nil {
			queue = append(queue, &Node{root1.Val.Right, root1.depth + 1})
		}
	}
	return num

}
