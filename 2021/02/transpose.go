package main

//867. 转置矩阵

//求给定矩阵的转置矩阵
//开辟空间
func transpose(matrix [][]int) [][]int {
    res := make([][]int, 0)
    for i := 0; i < len(matrix[0]); i++ {
        tmp := make([]int, 0)
        for j := 0; j < len(matrix); j++ {
            tmp = append(tmp, matrix[j][i])
        }
        res = append(res, tmp)
    }
    return res
}