package main

//802. 找到最终的安全状态

//dfs, 染色法
func eventualSafeNodes(graph [][]int) []int {
	color := make([]int, len(graph))
	res := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		if dfs(i, color, graph) {
			res = append(res, i)
		}
	}
	return res
}

// colors: WHITE 0, GRAY 1, BLACK 2;
// 白色表示该节点还没有被访问过
// 灰色表示该节点在栈中（这一轮搜索中被访问过）或者在环中
// 黑色表示该节点的所有相连的节点都被访问过，且该节点不在环中。
func dfs(node int, color []int, graph [][]int) bool {
	//递归出口
	if color[node] > 0 {
		return color[node] == 2
	}

	//标记在该轮中访问过
	color[node] = 1

	//遍历其他
	for _, v := range graph[node] {
		if color[node] == 2 {
			continue
		}
		if color[v] == 1 || !dfs(v, color, graph) {
			return false
		}
	}
	color[node] = 2
	return true
}
