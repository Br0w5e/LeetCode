//判断矩阵每条对角线上的元素是否相等
//弄清对角线元素关系
func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i := 0; i < n-1; i++{
        for j := 0; j < m-1; j++{
            if matrix[i][j] != matrix[i+1][j+1] {
                return false
            }
        }
    }
    return true
}