package main
// 根到叶路径上的不足节点
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 两遍遍历
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	flag := DFS(root, 0, limit)
	if flag {
		return root
	}
	return nil
}


func DFS (node *TreeNode, current int, limit int ) bool{
	if node == nil {
		return false
	}
	step := current + (*node).Val
	// fmt.Println("====", (*node).Val)

	// 叶节点
	if (*node).Left == nil && (*node).Right == nil {
		if step < limit {
			return false
		} else {
			return true
		}
	}
	left := DFS((*node).Left, current + (*node).Val, limit)
	right := DFS((*node).Right, current + (*node).Val, limit)
	// 左节点不满足，则删除
	if !left {
		(*node).Left = nil
	}
	// 右节点不满足，则删除
	if !right {
		(*node).Right = nil
	}
	return left || right
}
