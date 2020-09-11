package main
//797  0 到 n-1 的路径并输出
var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0)
	//源节点为0
	dfs(graph, []int{0})
	return res
}

func dfs(graph [][]int, path []int)  {
	tmp := make([]int, len(path))
	copy(tmp, path)
	//最后一个节点是不是终止节点
	curNode := tmp[len(tmp)-1]
	// 是最后一个节点,递归出口
	if curNode == len(graph)-1 {
		res = append(res, tmp)
		return
	}
	//不是最后一个节点，以该节点为源节点继续进行dfs遍历
	nodes := graph[curNode]
	for _, node := range nodes {
		dfs(graph, append(tmp, node))
	}
}