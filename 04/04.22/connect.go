//填充二叉树每一个节点的下一个指针
//层次填充
func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

//模版dfs
func connect(root *Node) *Node {
    if root == nil {
		return root
	}
	dfs(root, nil)
	return root
}

func dfs(node *Node, uncleNode *Node) {
	if node.Left == nil {
		return
	}
	node.Left.Next = node.Right
	if uncleNode != nil {
		node.Right.Next = uncleNode.Left
	}
	dfs(node.Left, node.Right)
	if node.Next != nil {
		dfs(node.Right, node.Next.Left)
	} else {
		dfs(node.Right, nil)
	}
}

//bfs
func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    queue := []*Node{root}
    for (len(queue) != 0) {
        count := len(queue)
        if queue[0].Left != nil {
            queue = append(queue, queue[0].Left)
        }
        if queue[0].Right != nil {
            queue = append(queue, queue[0].Right)
        }
        for i := 1; i < count; i++ {
            queue[i-1].Next = queue[i]
            if queue[0].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        queue = queue[count:]
    }
    return root
}
