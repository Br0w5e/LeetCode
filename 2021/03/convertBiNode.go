package main

//面试题 17.12. BiNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var s []*TreeNode

func in(root *TreeNode) {
	if root == nil {
		return
	}
	node := root
	in(node.Left)
	s = append(s, node)
	in(node.Right)
}

func buildByQueue() *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return build()
}

func build() *TreeNode {
	if len(s) == 0 {
		return nil
	}
	node := s[0]
	s = s[1:]
	node.Left = nil
	node.Right = build()
	return node
}

func convertBiNode(root *TreeNode) *TreeNode {
	//首先先中序遍历得到从小到大的有序队列
	//队列中第一个为头节点，所有节点做节点为空，右节点为下一个节点
	in(root)

	return buildByQueue()
}
