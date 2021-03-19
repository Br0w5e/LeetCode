package main

//59. 螺旋矩阵 II

//逐个操作
func generateMatrix(n int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	u, d, l, r := 0, n-1, 0, n-1
	step := 1
	for step <= n*n {
		for i := l; i <= r; i++ {
			res[u][i] = step
			step++
		}
		u++

		for i := u; i <= d; i++ {
			res[i][r] = step
			step++
		}
		r--

		for i := r; i >= l; i-- {
			res[d][i] = step
			step++
		}
		d--

		for i := d; i >= u; i-- {
			res[i][l] = step
			step++
		}
		l++
	}
	return res
}

//带方向控制
func generateMatrix(n int) [][]int {
	var dx = [4]int{-1, 0, 1, 0}
	var dy = [4]int{0, 1, 0, -1}
	if n == 0 {
		return [][]int{}
	}
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	for x, y, i, d := 0, 0, 0, 1; i < n*n; i++ {
		res[x][y] = i + 1
		a, b := x+dx[d], y+dy[d]
		if a < 0 || a == n || b < 0 || b == n || res[a][b] != 0 {
			d = (d + 1) % 4
			a, b = x+dx[d], y+dy[d]
		}
		x, y = a, b
	}
	return res
}
