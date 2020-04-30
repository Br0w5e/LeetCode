//计算投影的面积
//方法：求每个方向上的最大值
func projectionArea(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    up := 0
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if grid[i][j] != 0 {
                up++
            }
        }
    }
    left := 0
    for i := 0; i < n; i++{
        j, k := 0, m-1
        for j < k {
            if grid[i][j] < grid[i][k] {
                j++
            } else {
                k--
            }
        }
        left += grid[i][j]
    }
    front := 0 
    for i := 0; i < m; i++{
        j, k := 0, n-1
        for j < k {
            if grid[j][i] < grid[k][i] {
                j++
            } else {
                k--
            }
        }
        front += grid[j][i]
    }
    return up+left+front
}