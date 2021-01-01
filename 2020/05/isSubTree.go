//参考isSubStructure.go
package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	var res bool
	if s.Val == t.Val {
		res = judge(s, t)
	}
	if !res {
		res = isSubtree(s.Left, t)
	}
	if !res {
		res = isSubtree(s.Right, t)
	}
	return res
}

func judge(A *TreeNode, B *TreeNode)  bool{
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil{
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return judge(A.Left, B.Left) && judge(A.Right, B.Right)
}

func isSubtree2(s *TreeNode, t * TreeNode) bool {
	if s == nil {
		return false
	}
	return judge(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}