//计算可围成岛屿的最大面积
//思路：深度优先遍历，带置0
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j, grid))
			}
		}
	}
	return ans
}

func dfs(i, j int, grid[][]int) int{
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	num := 1
	grid[i][j] = 0
    for k := 0; k < 4; k++{
        num += dfs(i+dx[k], j+dy[k], grid)
    }

	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


//计算岛屿的周长
//思路：逐个遍历关注以前遍历过的对象
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++{
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1{
					res -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}

//计算岛屿的数量。
//思路 dfs
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}
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
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for k := 0; k < 4; k++{
		dfs(grid, i+dx[k], j+dy[k])
	}
}