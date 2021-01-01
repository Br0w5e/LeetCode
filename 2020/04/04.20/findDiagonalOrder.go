//按对角线顺序访问
func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    r, c := 0, 0
    row, col := len(matrix), len(matrix[0])
    res := make([]int, row*col)
    for i := 0; i < len(res); i++{
        res[i] = matrix[r][c]
        //r+c为遍历的层数，偶数向上，奇数向下
        if (r+c) % 2 == 0 {
            //往右边移动一格准备向下遍历
            if c == col-1 {
                r++
            } else if r == 0 {
                //向下移动一格准备向下遍历
                c++
            } else {
                //向上移动
                r--
                c++
            }
        } else {
            if  r == row - 1 {
                // 向右移动一格准备向上遍历
                c++
            } else if c == 0 {
                //向上移动一格准备向上遍历
                r++
            } else {
                r++
                c--
            }
        }
    }
    return res
}