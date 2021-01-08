package main
//399. 除法求值

//Floyd 算法 推荐

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//编号
	id := make(map[string]int)
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}
	//建立带权图
	graph := make([][]float64, len(id))
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]float64, len(id))
	}
	for i, eq := range  equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1/values[i]
	}

	//Folyd 算法，进行合并
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] > 0 && graph[k][j] > 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}
	//结果
	res := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasS || !hasE || graph[start][end] == 0 {
			res[i] = -1.0
		} else {
			res[i] = graph[start][end]
		}
	}
	return res
}

//bfs
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//给方程组中的每个变量编号
	id := make(map[string]int, 0)
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	//建图
	type edge struct {
		to int
		weight float64
	}

	graph := make([][]edge, len(id))
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v, 1/values[i]})
	}

	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == end {
				return ratios[v]
			}
			for _, e := range graph[v] {
				if w := e.to; ratios[w] == 0 {
					ratios[w] = ratios[v] * e.weight
					queue = append(queue, w)
				}
			}
		}
		return -1
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasS || !hasE {
			res[i] = -1
		} else {
			res[i] = bfs(start, end)
		}
	}
	return res
}

