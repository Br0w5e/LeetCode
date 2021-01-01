package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 创建入度表、邻接表
	indgree := make([]int, numCourses)
	adj := make([][]int, numCourses)

	for _, pre := range prerequisites {
		indgree[pre[0]]++
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}

	queue := make([]int, 0)
	for i, d := range indgree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		numCourses--
		for _, a := range adj[cur] {
			indgree[a]--
			if indgree[a] == 0 {
				queue = append(queue, a)
			}
		}
	}

	return numCourses == 0
}
