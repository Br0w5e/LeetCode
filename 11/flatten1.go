package main

//430. 扁平化多级双向链表
type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
//迭代
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	cur := root
	for cur != nil {
		after := cur.Next
		if cur.Child != nil {
			//扁平化
			child := cur.Child
			cur.Next = child
			child.Prev = cur
			cur.Child = nil
			//遍历child直到最后一个
			for child.Next != nil {
				child = child.Next
			}
			//最后一个child的前一个
			child.Next = after
			if after != nil {
				after.Prev = child
			}
		}
		cur = cur.Next
	}
	return root
}

//child当做右子树，Next当做左子树，先序遍历输出

