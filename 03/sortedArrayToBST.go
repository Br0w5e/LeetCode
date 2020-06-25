//将有序数组恢复成二叉排序树
func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums)	
}
func buildTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	res := &TreeNode{
		Val: nums[mid],
	}
	res.Left = buildTree(nums[:mid])
	res.Right = buildTree(nums[mid+1:])

	return res
}