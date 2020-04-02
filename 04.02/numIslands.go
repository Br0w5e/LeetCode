func numIslands(grid [][]byte) int {

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	count := 0
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == '1'{
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int){
    var dx = [4]int{-1, 1, 0, 0}
    var dy = [4]int{0, 0, -1, 1}
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for k := 0; k < 4; k++{
		dfs(grid, i+dx[k], j+dy[k])
	}
}