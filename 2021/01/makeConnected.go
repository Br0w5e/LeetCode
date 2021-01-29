package main

//1319. 连通网络的操作次数

//计算图的连通分支
func makeConnected(n int, connections [][]int) int {
	//结果
	res := 0

	//使n个节点连通最少需要n-1条边
	if len(connections) < n-1 {
		return -1
	}

	//生成图矩阵
	graph := make([][]int, n)
	for _, c := range connections {
		w, v := c[0], c[1]
		graph[w] = append(graph[w], v)
		graph[v] = append(graph[v], w)
	}

	//标记访问节点
	visited := make([]bool, n)

	//深度优先遍历，计算图的连通分支
	for i, v := range visited {
		if !v {
			res++
			dfs(&visited, graph, i)
		}
	}

	//最少需要移动图的连通分支减一条边
	return res-1
}

func dfs(visited *[]bool, graph [][]int, from int) {
	(*visited)[from] = true
	for _, to := range graph[from] {
		if !(*visited)[to] {
			dfs(visited, graph, to)
		}
	}
}
