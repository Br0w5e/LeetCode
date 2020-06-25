// 计算不规则积木体的表面积
// 方法一：逐个计算，加法
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//优化算法：减法
func surfaceArea(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 0 {
				//计算所以面
				res += 4 * grid[i][j] + 2
			}
			//减去多余面
			if j > 0 {
				res -= min(grid[i][j], grid[i][j-1])*2
			}
			if i > 0 {
				res -= min(grid[i][j], grid[i-1][j])*2
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}