package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
//dfs
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := visited[node]; ok {
		return v
	}

	newNode := &Node{
		Val: node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode

	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}

//bfs

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := map[*Node]*Node{}
	queue := []*Node{node}
	visited[node] = &Node{node.Val, []*Node{}}
	//常规操作
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, neighbor := range n.Neighbors {
			//未访问节点
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{neighbor.Val, []*Node{}}
				queue = append(queue, neighbor)
			}
			//更新邻居节点
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}