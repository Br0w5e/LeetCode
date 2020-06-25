//子树判断
//判断是不是子树
//参考isSubTree.go
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var res bool
	if A.Val == B.Val {
		res = judge(A, B)
	}

	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)
	}

	return res

}

func judge(A *TreeNode, B *TreeNode) bool{
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val != B.Val{
		return false
	}

	return judge(A.Left, B.Left) && judge(A.Right, B.Right)
}

func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil{
		return true
	}
	if A == nil || B == nil {
		return false
	}
	return judge(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
