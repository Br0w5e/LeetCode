package main

//5702. 找出星型图的中心节点
func findCenter(edges [][]int) int {
	m := make(map[int]int)
	n := len(edges)
	for _, edge := range edges {
		m[edge[0]]++
		m[edge[1]]++
		if m[edge[0]] == n {
			return edge[0]
		}
		if m[edge[1]] == n {
			return edge[1]
		}
	}
	return -1
}
