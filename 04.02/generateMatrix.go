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