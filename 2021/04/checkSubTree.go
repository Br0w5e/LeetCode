package main

//参考一下，isSame
//面试题 04.10. 检查子树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//这个代码使用了判断子结构的
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return checkSubTree(t1.Left, t2.Left) && checkSubTree(t1.Right, t2.Right)
	} else {
		return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
	}
}


//这个才是真正的代码
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return false
	}
	if t1.Val == t2.Val {
		if isSame(t1, t2) {
			return true
		}
	}
	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func isSame(t1, t2 *TreeNode) bool {
	//必须要为子树，不能是子结构
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}
