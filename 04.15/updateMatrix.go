//设起始 level = 0
//遍历矩阵遇到0就把4周的1设置为 level - 1
//level--
//继续第二次遍历 直到没有新level
func updateMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0 {
        return nil
    }
    n, m := len(matrix), len(matrix[0])
    var helper func(i int, j int)
    level := 0
    helper = func(i int, j int) {
        if i < 0 || i >= n || j < 0 || j >= m {
            return
        }
        if matrix[i][j] == 1 {
            matrix[i][j] = level-1
        }
    }
    //直到flag = false
    flag := true
    for flag {
        flag = false
        for i := 0; i < n; i++{
            for j := 0; j < m; j++{
                if matrix[i][j] == level {
                    flag = true
                    helper(i-1, j)
                    helper(i+1, j)
                    helper(i, j+1)
                    helper(i, j-1)
                }
            }
        }
        level--
    }
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            matrix[i][j] = -matrix[i][j]
        }
    }
    return matrix
}