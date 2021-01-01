package main

//1609. 奇偶树
//方法一：层次遍历，然后判断结果

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		res = append(res, make([]int, 0))
		n := len(queue)
		for i := 0; i < n; i++ {
			top :=  queue[0]
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
	for i, level := range res {
		if i % 2 == 0{
			for j := 0; j < len(level)-1; j++ {
				if level[j] >= level[j+1] || level[j]%2 == 0{
					return false
				}
			}
			if level[len(level)-1]%2 == 0 {
				return false
			}
		} else {
			for j := 0; j < len(level)-1; j++ {
				if level[j] <= level[j+1] || level[j]%2 == 1 {
					return false
				}
			}
			if level[len(level)-1]%2 == 1 {
				return false
			}
		}
	}
	return true
}


//方法2：边遍历边判断
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//层次遍历
func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	ans := 1
	for len(queue) > 0 {
		//存放最后一个节点值
		next := 0
		n := len(queue)
		//奇数层
		if ans%2 == 1 {
			//遍历完奇数层的所有节点
			for i := 0; i < n; i++ {
				//判断逻辑
				top :=  queue[0]
				if (i > 0 && top.Val <= next) || top.Val % 2 == 0 {
					return false
				}
				next = top.Val
				//出队
				queue = queue[1:]
				//子节点加入队列
				if top.Left != nil {
					queue = append(queue, top.Left)
				}
				if top.Right != nil {
					queue = append(queue, top.Right)
				}
			}
		} else {
			//遍历完偶数层的所有节点
			for i := 0; i < n; i++ {
				//判断逻辑
				top :=  queue[0]
				if (i > 0 && top.Val >= next) || top.Val % 2 == 1{
					return false
				}
				next = top.Val
				queue = queue[1:]
				//子节点加入队列
				if top.Left != nil {
					queue = append(queue, top.Left)
				}
				if top.Right != nil {
					queue = append(queue, top.Right)
				}
			}
		}
		//层次加一
		ans++
	}
	return true
}