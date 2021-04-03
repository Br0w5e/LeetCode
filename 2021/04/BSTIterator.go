package main

//173. 二叉搜索树迭代器
import (
	"strconv"
	"strings"
)

//二叉搜索树迭代树
//等价于二叉树的中序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
	}
}

//bfs
func (this *BSTIterator) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	//取栈顶元素
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	//返回值
	val := this.cur.Val
	//右边节点
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

//dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
	var bst BSTIterator
	bst.inorder(root)
	return bst
}

func (this *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	this.inorder(node.Left)
	this.nums = append(this.nums, node.Val)
	this.inorder(node.Right)
}

func (this *BSTIterator) Next() int {
	val := this.nums[0]
	this.nums = this.nums[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
