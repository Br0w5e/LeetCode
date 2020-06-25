//判断二叉树是否有指定和的路径
//对着pathSum来看
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil{
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right,sum)
}