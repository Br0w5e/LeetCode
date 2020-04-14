//将两个链表相加
//方法1：栈操作
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	flag, sum := 0, 0
	var res *ListNode
	res = nil 
	for len(s1) > 0 || len(s2) > 0 {
		a, b := 0, 0
		if len(s1) > 0 {
			//最低位
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		sum = a+b+flag
		flag = sum/10
		sum %= 10
		cur := &ListNode{
			Val : sum,
			Next : res,
		}
		res = cur
	}
	if flag > 0 {
		cur := &ListNode {
			Val : flag,
			Next : res,
		}
		res = cur
	}
	return res
}

//方法2：翻转链表然后求和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := reverse(l1)
	h2 := reverse(l2)
	var temp *ListNode
	var newHead *ListNode
	up := 0
	for {
		if h1 == nil && h2 == nil && up == 0 {
			break
		}
		newNode, newUp := add(h1, h2, up)
		up = newUp
		if temp != nil {
			temp.Next = newNode
		} else {
			newHead = newNode
		}
		temp = newNode
		if h1 != nil {
			h1 = h1.Next
		}
		if h2 != nil {
			h2 = h2.Next
		}
	}
	return reverse(newHead)
}

func add(l1 *ListNode, l2 *ListNode, up int) (*ListNode, int) {
	var res int
	var newUp int
	if l1 == nil && l2 == nil {
		res = up
		newUp = 0
	} else if l1 == nil {
		res = (l2.Val + up) % 10
		newUp = (l2.Val + up) / 10
	} else if l2 == nil {
		res = (l1.Val + up) % 10
		newUp = (l1.Val + up) / 10
	} else {
		res = (l1.Val + l2.Val + up) % 10
		newUp = (l1.Val + l2.Val + up) / 10
	}
	return &ListNode{res, nil}, newUp
}

func reverse(head *ListNode) *ListNode {
	var front *ListNode
	var next *ListNode
	for {
		if head == nil {
			break
		}
		next = head.Next
		head.Next = front
		front = head
		head = next
	}
	return front
}
