package main

//94  二叉树的中序遍历

//dfs 应用
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderDfs(&res, root)
	return res
}

func inorderDfs(res *[]int, root *TreeNode) {
	// if root == nil {
	// 	return
	// }
	// dfs(res, root.Left)
	// *res = append(*res, root.Val)
	// dfs(res, root.Right)
	if root != nil {
		inorderDfs(res, root.Left)
		*res = append(*res, root.Val)
		inorderDfs(res, root.Right)
	}
}


//栈应用

func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	curNode := root

	for len(stack) > 0 || curNode != nil{
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curNode.Val)
		curNode = curNode.Right
	}
	return res
}