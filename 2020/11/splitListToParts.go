package main

//725. 分隔链表

func splitListToParts(root *ListNode, k int) []*ListNode {
	p := root
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	avgeLength := length/k
	reminder := length%k

	res := make([]*ListNode, k)
	head := root
	var pre *ListNode
	for i := 0; i < k; i++ {
		res[i] = head
		tmpLength := 0
		//当前长度
		if reminder > 0 {
			tmpLength = avgeLength+1
		} else {
			tmpLength = avgeLength
		}

		//处理节点
		for j := 0; j < tmpLength; j++ {
			pre = head
			head = head.Next
		}

		if pre != nil {
			pre.Next = nil
		}

		//处理余数
		if reminder > 0 {
			reminder--
		}
	}
	return res
}
