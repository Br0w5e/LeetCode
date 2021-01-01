package main

// 107 二叉树的层次遍历
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		n := len(queue)
		res = append(res, make([]int, 0))
		for i := 0; i < n; i++ {
			top := queue[0]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			res[ans] = append(res[ans], top.Val)
			queue = queue[1:]
		}
		ans++
	}

	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}
