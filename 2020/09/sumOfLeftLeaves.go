package main

//404. 左叶子之和
//dfs
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	//左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
		dfs(root.Left, res)
	}
	//左边但不是叶子节点
	if root.Left != nil {
		dfs(root.Left, res)
	}
	//右边
	if root.Right != nil {
		dfs(root.Right, res)
	}
}

//stack
//逐个左子树清算
func sumOfLeftLeaves2(root *TreeNode) int {
	res := 0
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			if root != nil && root.Left == nil && root.Right == nil {
				res += root.Val
			}
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	}
	return res
}
