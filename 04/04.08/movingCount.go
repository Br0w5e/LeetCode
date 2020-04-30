//确定机器人的移动范围
//设置辅助数组，进行判断
var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}
func movingCount(m int, n int, k int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++{
		grid[i] = make([]int, n)
	}
	return dfs(0, 0, m, n, k, grid)
}
func dfs(x int, y int, m int, n int, k int, grid [][]int) int {
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 1 || digitSum(x, y) > k {
		return 0
	}
	grid[x][y] = 1 
	sum := 1
	for i := 0; i < 4; i++{
		sum += dfs(x+dx[i], y+dy[i], m, n, k, grid)
	}
	return sum
}
func digitSum(num1, num2 int) int {
    res := 0 
    for num1 != 0 {
        res += num1%10
        num1 /= 10
    }
    for num2 != 0 {
        res += num2%10
        num2 /= 10
    }
    return res
}