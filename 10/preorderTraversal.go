package main

//二叉树的先序遍历
//dfs
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	//递归出口
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

//迭代
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		//此过程中默认存在遍历根节点和左子树节点，只需要操作遍历右子树即可
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}