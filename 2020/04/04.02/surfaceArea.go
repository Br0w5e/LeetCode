func surfaceArea(grid [][]int) int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	n := len(grid)
	ans := 0
	for x := 0; x < n; x++{
		for y := 0; y < n; y++{
			if grid[x][y] > 0 {
				ans += 2 //上下底面
				for k := 0; k < 4; k++{
					nx := x + dx[k]
					ny := y + dy[k]
					nv := 0
					if nx >= 0 && nx < n && ny >= 0 && ny < n  {
						nv = grid[nx][ny]
					}
					//侧面
					ans += max(grid[x][y]-nv, 0)
				}
			}
		}
	}
	return ans
}