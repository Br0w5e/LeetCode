package main

//589. N 叉树的前序遍历

//非递归

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//模仿二叉树
//非递归

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children)-1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}


//递归
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, node := range root.Children {
		dfs(node, res)
	}
}