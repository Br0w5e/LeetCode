package main

//222. 完全二叉树的节点个数

func countNodes(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	*res++
	dfs(root.Left, res)
	dfs(root.Right, res)
}

//dfs
func countNodes1(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		res++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return res
}