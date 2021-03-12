package main

//690. 员工的重要性

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

//DFS 或者 BFS
func getImportance(employees []*Employee, id int) int {
	res := 0
	m := make(map[int]*Employee)
	for _, v := range employees {
		m[v.Id] = v
	}
	queue := make([]int, 0)
	queue = append(queue, id)
	for len(queue) != 0 {
		cur := m[queue[0]]
		queue = queue[1:]
		res += cur.Importance
		for i := 0; i < len(cur.Subordinates); i++ {
			queue = append(queue, cur.Subordinates[i])
		}
		//queue = append(queue, cur.Subordinates...)
	}
	return res
}

//dfs
func getImportance1(employees []*Employee, id int) int {
	res := 0
	m := make(map[int]*Employee)
	for _, v := range employees {
		m[v.Id] = v
	}

	var dfs func(Id int)
	dfs = func(Id int) {
		cur := m[Id]
		res += cur.Importance
		for _, v := range cur.Subordinates {
			dfs(v)
		}
	}

	dfs(id)
	return res
}