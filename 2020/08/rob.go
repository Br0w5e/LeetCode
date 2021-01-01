package main

func rob(root *TreeNode) int {
	mp := make(map[*TreeNode]int)
	return robImpl(root, mp)
}

func robImpl(root *TreeNode, mp map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if val, ok := mp[root]; ok {
		return val
	}
	money := root.Val
	if root.Left != nil {
		money += robImpl(root.Left.Left, mp) + robImpl(root.Left.Right, mp)
	}
	if root.Right != nil {
		money += robImpl(root.Right.Left, mp) + robImpl(root.Right.Right, mp)
	}
	r := max(money, robImpl(root.Left, mp) + robImpl(root.Right, mp))
	mp[root] = r
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
