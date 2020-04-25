//二叉搜索树迭代树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
	numList   []int
	nextIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	Iterator := BSTIterator{
		numList:   make([]int, 0),
		nextIndex: 0,
	}
	inorder(root, &Iterator.numList)

	return Iterator
}

func inorder(root *TreeNode, numList *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, numList)
	*numList = append(*numList, root.Val)
	inorder(root.Right, numList)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.nextIndex > len(this.numList)-1 {
		return -1
	}
	res := this.numList[this.nextIndex]
	this.nextIndex++
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.nextIndex > len(this.numList)-1 {
		return false
	} else {
		return true
	}
}