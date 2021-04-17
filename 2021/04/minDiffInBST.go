package main
//783. 二叉搜索树节点最小距离
import "math"

//dfs
func minDiffInBST(root *TreeNode) int {
	nums := make([]int, 0)
	dfs(root, &nums)
	res := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] < res {
			res = nums[i+1]-nums[i]
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


//bfs，树的先序遍历
func minDiffInBST(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	nums := make([]int, 0)
	for len(stack) > 0 || root != nil {
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
		if nums[i] - nums[i-1] < res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}