//返回二叉树中所有满足指定和的路径
package main

//寻找路径和，保证跟节点到叶子节点
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root, sum, []int{}, &ret)
	return ret
}

func dfs(root *TreeNode, sum int, arr []int, ret *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*ret = append(*ret, tmp)
	}
	dfs(root.Left, sum-root.Val, arr, ret)
	dfs(root.Right, sum-root.Val, arr, ret)
	arr = arr[:len(arr)-1]
}

//面试题 04.12. 求和路径
//不用保证根节点到 叶子节点
func pathSum(root *TreeNode, sum int) int {
	ret := 0
	solve(root, sum, &ret)
	return ret
}

func solve(root *TreeNode, sum int, ret *int) {
	if root == nil {
		return
	}
	dfs(root, sum, ret)
	solve(root.Left, sum, ret)
	solve(root.Right, sum, ret)
}

func dfs(root *TreeNode, sum int, ret *int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		*ret++
	}
	dfs(root.Left, sum-root.Val, ret)
	dfs(root.Right, sum-root.Val, ret)
}
