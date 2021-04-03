package main

//面试题 04.06. 后继者

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//迭代法
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	res := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		//回溯
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node)
		node = node.Right
	}

	for i, n := range res {
		if n == p && i != len(res)-1 {
			return res[i+1]
		}
	}
	return nil
}

//递归发
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	res := make([]*TreeNode, 0)
	dfs(root, &res)
	for i, node := range res {
		if node == p && i != len(res)-1 {
			return res[i+1]
		}
	}
	return nil
}

func dfs(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root)
	dfs(root.Right, res)
}

//二叉搜索树
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	//二叉搜索树性质，在右边搜索
	if p.Val >= root.Val {
		return inorderSuccessor(root.Right, p)
	}

	if left := inorderSuccessor(root.Left, p); left == nil {
		return root
	} else {
		return left
	}
}
