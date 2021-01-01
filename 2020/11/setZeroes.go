package main
//面试题 01.08. 零矩阵
//将矩阵0所在的行和列都置位0
//方法：求出0所在行和列的行列号，然后进行求解
func setZeroes(matrix [][]int)  {
    n, m := len(matrix), len(matrix[0])
    grid := make([][]int, 0)
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if matrix[i][j] == 0 {
                tmp := make([]int, 2)
                tmp[0], tmp[1] = i, j
                grid = append(grid, tmp)
            }
        }
    }
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            for k := 0; k < len(grid); k++{
                if  i == grid[k][0] || j == grid[k][1] {
                    matrix[i][j] = 0
                }
            }
        }
    }
    return
}

//map标记
func setZeroes1(matrix [][]int)  {
    m := make([][]int, 0)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                m = append(m, []int{i, j})
            }
        }
    }
    for _, v := range m {
        i := v[0]
        j := v[1]
        for k := 0; k < len(matrix[i]); k++ {
            //行置0
            matrix[i][k] = 0
        }

        for k := 0; k < len(matrix); k++ {
            //列置0
            matrix[k][j] = 0
        }
    }
}