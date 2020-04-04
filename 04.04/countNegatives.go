//统计有序矩阵中的负数个数
func countNegatives(grid [][]int) int {
    res := 0
    for _, row := range grid {
        for _, num := range row {
            if num < 0 {
                res++
            }
        }
    }
    return res
}

//方法二
func countNegatives(grid [][]int) int {
    res := 0
    for i := 0; i < len(grid); i++{
        for j := len(grid[0])-1; j >= 0; j-- {
            if grid[i][j] < 0{
                res++
            } else {
                break
            }
        }
    }
    return res
}