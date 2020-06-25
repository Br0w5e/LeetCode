// 二叉树的三种遍历方法：递归
// 前序遍历
var preres []int
func preorderTraversal(root *TreeNode) []int {
	preres = make([]int, 0)
	dfs(root)
	return preres
}

func dfs(root *TreeNode){
	if root != nil {
		preres = append(preres, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}

//中序遍历
var inres []int
func inorderTraversal(root *TreeNode) []int {
	inres = make([]int, 0)
	dfs(root)
	return inres
}
func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		inres = append(inres, root.Val)
		dfs(root.Right)
	}
}

//后序遍历
var postres []int 
func postorderTraversal(root *TreeNode) []int {
	postres = make([]int, 0)
	dfs(root)
	return postres
}
func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		dfs(root.Right)
		postres = append(postres, root.Val)
	}
}