package main

//可以到达所有点点最小点数

//统计终点
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	tmp := make([]int, n)
	//标记
	for _, edge := range edges {
		tmp[edge[1]] += 1
	}
	//寻找没有终点不在的
	res := make([]int, 0)
	for i, num := range tmp {
		if num == 0 {
			res = append(res, i)
		}
	}
	return res
}
