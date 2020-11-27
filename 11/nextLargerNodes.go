package main


//1019. 链表中的下一个更大节点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//数组
func nextLargerNodes(head *ListNode) []int {
	tmp := make([]int, 0)
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	res := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		for j := i+1; j < len(tmp); j++ {
			if tmp[j] > tmp[i] {
				res[i] = tmp[j]
				break
			}
		}
	}
	return res
}

func nextLargerNodes1(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		cur := head
		tmp := cur.Val
		max := cur.Val
		for cur != nil {
			if cur.Val > max {
				max = cur.Val
				break
			}
			cur = cur.Next
		}
		if tmp == max {
			max = 0
		}
		res = append(res, max)
		head = head.Next
	}
	return res
}