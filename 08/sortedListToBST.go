//将排序链表转化为二叉搜索树
//思路一：先将链表转化为数组，再转化为二叉搜索树
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
package main
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := getVal(head)
	return buildTree(nums)
}

func getVal(head *ListNode) []int{
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func buildTree(nums []int) *TreeNode{
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	res := &TreeNode{
		Val : nums[mid],
	}
	res.Left = buildTree(nums[:mid])
	res.Right = buildTree(nums[mid+1:])
	return res
}


//思路二：直接转化
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	last, mid := findMiddle(head)
	if last.Next == mid {
		last.Next = nil
	} else {
		head = nil
	}

	res := &TreeNode{
		Val: mid.Val,
	}
	res.Left = sortedListToBST(head)
	res.Right = sortedListToBST(mid.Next)

	return res
}

func findMiddle(head *ListNode) (*ListNode, *ListNode)  {
	//一步节点
	one := head
	//两步节点
	two := head
	last := two
	for two != nil{
		two = two.Next
		if two != nil {
			two = two.Next
		} else {
			break;
		}
		last = one
		one = one.Next
	}
	return last, one
}