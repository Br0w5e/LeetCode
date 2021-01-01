package main

//530. 二叉搜索树的最小绝对差
//783. 二叉搜索树节点最小距离

//理解题目中的，绝对值和二叉搜索树的关系
func getMinimumDifference(root *TreeNode) int {
	nums := make([]int, 0)
	//二叉搜索树的中序遍历为升序排列
	dfs(root, &nums)
	res := nums[1]-nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i-1] < res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

func getMinimumDifference1(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	nums := make([]int, 0)
	//中序遍历
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}
	//求解结果
	res := nums[1]-nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i-1] < res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func minDiffInBST(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	nums := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	res := nums[1]-nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] < res{
			res = nums[i]-nums[i-1]
		}
	}
	return res
}