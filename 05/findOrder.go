package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result []int
		invalid bool
		dfs func(u int)
	)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if invalid {
					return
				}
			} else if visited[v] == 1 {
				invalid = true
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	for i := 0; i < numCourses && !invalid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if invalid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	ingress := make([]int, numCourses)
	var stack []int
	var edge = make([][]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		ingress[prerequisites[i][0]]++
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]],  prerequisites[i][0])
	}

	for i := 0; i < numCourses; i++ {
		if ingress[i] == 0 {
			stack = append(stack, i)
		}
	}
	result := BFS(stack, ingress, edge)
	if len(result) != numCourses {
		return nil
	}

	return result
}

func BFS(stack []int, ingress []int, edge [][]int) []int {
	var result []int
	for len(stack) > 0 {
		cenCourse := stack[0]
		//fmt.Println("1:", stack, edge[cenCourse])
		for i := 0; i < len(edge[cenCourse]); i++ {
			nextCourse := edge[cenCourse][i]
			if nextCourse == cenCourse {
				continue
			}

			ingress[nextCourse]--
			if ingress[nextCourse] <= 0 {
				//fmt.Println("2:", stack, nextCourse)
				stack = append(stack, nextCourse)
			}
		}
		result = append(result, cenCourse)
		stack = stack[1:]
	}
	return result
}