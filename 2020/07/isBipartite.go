package main
//判断图是否为二分图
//方法：深度优先遍历，染色法。
//把某节点的相邻节点染成相同颜色，如果邻接节点颜色相同则不能二分
var (
	visited map[int]bool
	color map[int]int
)

func isBipartite(graph [][]int) bool {
	visited = make(map[int]bool)
	color = make(map[int]int)
	IsB := true

	for i, _ := range graph {
		if visited[i] != true {
			IsB = dfs(i, 0, graph)
			if IsB != true {
				return IsB
			}
		}
	}
	return IsB
}

func dfs(v int, c int, g [][]int) bool {
	//访问节点标记为true
	visited[v] = true
	items := g[v]
	color[v] = c

	for _, item := range items {
		//节点未访问
		if visited[item] != true {
			if !dfs(item, 1-c, g) {
				return false
			}
		} else {
			//节点已访问，且被染色
			if color[item] == c {
				return false
			}
		}
	}
	return true
}
