package main

import "strconv"

//257 二叉树的所有路径

//dfs
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	constructPaths(root, "")
	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths(root.Left, pathSB)
			constructPaths(root.Right, pathSB)
		}
	}
}

//bfs
func binaryTreePaths2(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{}
	pathQueue := []string{}
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}
