package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten1(root *TreeNode) {
	list := preOrderTraversal1(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preOrderTraversal1(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	for root != nil {
		res = append(res, root)
		res = append(res, preOrderTraversal1(root.Left)...)
		res = append(res, preOrderTraversal1(root.Right)...)
	}
	return res
}
//两种递归的方式不同
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []*TreeNode
func flatten(root *TreeNode) {
	preOrderTraversal(root)
	for i := 1; i < len(res); i++ {
		prev, curr := res[i-1], res[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preOrderTraversal(root *TreeNode) {
	//递归出口
	if root != nil {
		res = append(res, root)
		preOrderTraversal(root.Left)
		preOrderTraversal(root.Right)
	}
}


func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		//出栈
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right

		//栈的特点，FILO
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}