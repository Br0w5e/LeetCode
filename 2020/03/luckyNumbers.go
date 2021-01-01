//幸运数字，计算每一行行最小，列最大的数字称之为幸运数字
func luckyNumbers (matrix [][]int) []int {
    res := make([]int, 0)
    for i, n := 0, len(matrix); i < n; i++{
        for j, m := 0, len(matrix[0]); j < m; j++{
            flag := true
            for k := 0; flag == true && k < m; k++{
                if k == j {
                    continue
                }
                if matrix[i][j] > matrix[i][k]{
                    flag = false
                }
            }
            for k := 0; flag == true && k < n; k++{
                if k == i {
                    continue
                }
                if matrix[i][j] < matrix[k][j]{
                    flag = false
                }
            }
            if flag == true{
                res = append(res, matrix[i][j])
            }
        } 
    }
    return res
}

//寻找列最校和行最大的，并判断是不是同一元素
func luckyNumbers (matrix [][]int) []int {
    rows, cols := len(matrix), len(matrix[0])
    res := make([]int, 0)
    for i := 0; i < rows; i++{
        left, right := 0, cols-1
        for left < right{
            if matrix[i][left] > matrix[i][right] {
                left++
            } else {
                right--
            }
        }
        up, down := 0, rows-1
        for up < down{
            if matrix[up][left] > matrix[down][left]{
                down--
            } else {
                up++
            }
        }
        if i == up {
            res = append(res, matrix[up][left])
        }
    }
    return res
}