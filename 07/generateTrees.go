package main

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	//递归出口
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		//左子树
		leftTrees := helper(start, i-1)
		//右子树
		rightTrees := helper(i+1, end)
		for _, l := range leftTrees {
			for _, r := range rightTrees {
				root := &TreeNode{Val: i, Left: l, Right: r}
				res = append(res, root)
			}
		}
	}
	return res
}

func main()  {
	for _, val := range generateTrees(3) {
		fmt.Println(val.Val)
	}
}