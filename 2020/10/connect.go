//116. 填充每个节点的下一个右侧节点指针
//完美二叉树
//层次填充
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//bfs
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		count := len(queue)
		//左右子节点入队
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		//设置指针
		for i := 1; i < count; i++ {
			queue[i-1].Next = queue[i]
			if queue[0].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		//设置完出队
		queue = queue[count:]
	}
	return root
}

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

