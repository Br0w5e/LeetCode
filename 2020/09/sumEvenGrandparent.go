package main

//1315. 祖父节点值为偶数的节点和

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				res += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				res += root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				res += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				res += root.Right.Right.Val
			}
		}
	}
	res += sumEvenGrandparent(root.Left)
	res += sumEvenGrandparent(root.Right)
	return res
}
