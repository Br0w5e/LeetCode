package main

func longestIncreasingPath(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	cache := make([][]int, row)

	for i := 0; i < row; i++ {
		cache[i] = make([]int, col)
	}
	rst := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			rst = Max(rst, helper(matrix, i, j, row, col, cache))
		}
	}
	return rst
}

func helper(matrix [][]int, i, j, row, col int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	max := 1
	if i >= 1 && matrix[i][j] < matrix[i-1][j] {
		max = Max(max, 1+helper(matrix, i-1, j, row, col, cache))
	}
	if j >= 1 && matrix[i][j] < matrix[i][j-1] {
		max = Max(max, 1+helper(matrix, i, j-1, row, col, cache))
	}
	if i+1 < row && matrix[i][j] < matrix[i+1][j] {
		max = Max(max, 1+helper(matrix, i+1, j, row, col, cache))
	}
	if j+1 < col && matrix[i][j] < matrix[i][j+1] {
		max = Max(max, 1+helper(matrix, i, j+1, row, col, cache))
	}
	cache[i][j] = max
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
