//二叉搜索树中插入值
//方法：递归

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if val < root.Val {
		if root.Left == nil{
			root.Left = &TreeNode{Val : val}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val : val}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}