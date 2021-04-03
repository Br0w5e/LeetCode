package main

//LCP 07. 传递信息
func numWays(n int, relation [][]int, k int) int {
	res := 0
	for  i := 0; i < len(relation); i++ {
		if relation[i][0] == 0 {
			dfs(relation, relation[i][1], 1, k, n, &res)
		}
	}
	return res
}

func dfs(relation [][]int, start int, step int, k int, n int, res *int) {
	//步数等于k并且 到达n-1号
	if step == k {
		if start == n - 1 {
			*res++
		}
		return
	}
	for i := 0; i < len(relation); i++ {
		if relation[i][0] == start {
			dfs(relation, relation[i][1], step+1, k, n, res)
		}
	}
}
