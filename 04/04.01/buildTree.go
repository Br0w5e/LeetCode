//递归，了解二叉树的性质
// 回复二叉树， 由前序和中序回复
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	//这个返回值得掌握
	res := &TreeNode{
		Val: preorder[0],
	}

	if len(inorder) == 1 {
		return res
	}

	//重点
	idx := indexOf(res.Val, inorder)
	//在preorder中idx表示元素个数
	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:],inorder[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

//index得到方法
/*
func indexOf(val int, nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
*/

//由中序和后序回复 
func buildTree(inorder []int, postorder []int * TreeNode) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	res := &TreeNode{
		Val: postorder[n-1],
	}
	if n == 1{
		return res
	}

	//重点在 postorder中idx表元素个数
	idx := getIndex(inorder, res.Val)
	res.Left = buildTree(inorder[:idx], postorder[:idx])
	res.Right = buildTree(inorder[idx+1:], postorder[idx:n-1])
}

func getIndex(nums []int, val int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}