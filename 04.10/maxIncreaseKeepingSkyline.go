//求以原来矩阵行最大和列最大为值的新矩阵，使得新矩阵增量最大
//思路
/*
将新矩阵置位行最大矩阵
根据列最大值修正新矩阵
求新矩阵中每个元素和原来元素的差值，并求和返回
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    gridNew := make([][]int, n)
    max := getMax(grid, n, m)
    for i := 0; i < n; i++{
        gridNew[i] = make([]int, m)
    }
    //先设置为行最大
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            gridNew[i][j] = max[0][i]
        }
    }
    //修正为列最大
    for i := 0; i < m; i++{
        for j := 0; j < n; j++{
            if gridNew[j][i] > max[1][i]{
                gridNew[j][i] = max[1][i]
            }
        }
    }
    //求原来的差值
    res := 0
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            res += gridNew[i][j]-grid[i][j]
        }
    }
    return res

}

func getMax(grid [][]int, n int, m int) [][]int {
    res := make([][]int, 2)
    //行最大
    for _, value := range grid {
        left, right := 0, m-1
        for left < right {
            if value[left] < value[right] {
                left++
            } else {
                right--
            }
        }
        res[0] = append(res[0], value[left])
    }
    //列最大
    for i := 0; i < m; i++{
        up, down := 0, n-1
        for up < down {
            if grid[up][i] < grid[down][i] {
                up++
            }  else {
                down--
            }
        }
        res[1] = append(res[1], grid[up][i])
    }
    return res
}