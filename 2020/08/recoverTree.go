package main
// 	恢复二叉搜索树
func recoverTree(root *TreeNode)  {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapped(nums)
	recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree2(root *TreeNode)  {
	var (
		curNode    = root
		preNode    *TreeNode
		firstNode  *TreeNode
		secondNode *TreeNode
	)
	handle := func() {
		if preNode != nil && preNode.Val > curNode.Val {
			if firstNode == nil {
				firstNode = preNode
			}
			secondNode = curNode
		}
		preNode = curNode
		curNode = curNode.Right
	}
	for curNode != nil {
		if curNode.Left != nil {
			curNodePre := curNode.Left
			for curNodePre.Right != nil && curNodePre.Right != curNode {
				curNodePre = curNodePre.Right
			}
			if curNodePre.Right != curNode {
				curNodePre.Right = curNode
				curNode = curNode.Left
				continue
			}
			curNodePre.Right = nil
			handle()
		} else {
			handle()
		}
	}
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}