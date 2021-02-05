package main

//5665. 从相邻元素对还原数组

//dfs
//头尾两个仅出现一次，其他都出现两次， 使用map映射
func restoreArray(adjacentPairs [][]int) []int {
	m := make(map[int][]int, 0)

	for _, a := range adjacentPairs {
		m[a[0]] = append(m[a[0]], a[1])
		m[a[1]] = append(m[a[1]], a[0])
	}
	visited := make(map[int]bool, 0)
	ends := make([]int, 0)
	for k, v := range m {
		if len(v) != 2 {
			ends = append(ends, k)
		}
	}
	res := make([]int, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		res = append(res, cur)
		visited[cur] = true
		for _, n := range m[cur] {
			_, ok := visited[n]
			if !ok {
				visited[n] = true
				dfs(n)
			}
		}
	}
	dfs(ends[0])
	return res
}
