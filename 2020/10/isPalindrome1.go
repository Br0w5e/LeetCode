package main

//判断链表是不是回文链表
//将数值存储在数组中然后进行判断
func isPalindrome(head *ListNode) bool {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	left, right := 0, len(res)-1
	for left < right {
		if res[left] != res[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//遍历一半，然后翻转后半部分链表，接着遍历
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail := reverse(slow)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		tail = tail.Next
		head = head.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}