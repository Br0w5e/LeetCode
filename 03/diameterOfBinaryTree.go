//计算二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		result = math.MinInt32
	)
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, result)
	right := helper(root.Right, result)

	if left + right > *result {
		*result = left + right
	}
	return getMax(left, right) + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//示例代码
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	treeHeight(root, &max)

	return max
}

func treeHeight(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	leftH := treeHeight(root.Left, max)
	rightH := treeHeight(root.Right, max)

	*max = maxInt(rightH+leftH, *max)

	return maxInt(leftH, rightH) + 1
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}