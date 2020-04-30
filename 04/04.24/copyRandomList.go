//复制复杂链表
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 第一步复制节点 并插入到当前节点的后面
	cur := head
	for cur != nil {
		new_node := new(Node)
		new_node.Val = cur.Val
		new_node.Next = cur.Next
		cur.Next = new_node
		cur = new_node.Next
	}
	// 第二步 更新第一步中复制节点的 Random
	// 逻辑就是：由于复制的节点紧跟在原始节点的后面，所以复制节点的Random 就相当于 原始节点Random的后方（就是复制出的Random）
	cur = head
	for cur != nil {
		if cur.Random != nil { // 注意此处判断  cur.Random为nil的话 cur.Random.Next会报错
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 第三步：提取出新链表和旧链表 （理解为奇数为旧链表 偶数为新链表）
	cur_old_list := head
	cur_new_list := head.Next
	new_head := head.Next
	for cur_old_list != nil {
		cur_old_list.Next = cur_old_list.Next.Next
		if cur_new_list.Next != nil { // 注意此处  判断  cur_new_list.Next为nil的话，cur_new_list.Next.Next会报错
			cur_new_list.Next = cur_new_list.Next.Next
		}

		cur_old_list = cur_old_list.Next
		cur_new_list = cur_new_list.Next
	}

	return new_head

}