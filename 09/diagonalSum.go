package main

//1572. 矩阵对角线元素的和

func diagonalSum(mat [][]int) int {
	res := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				res += mat[i][j]
			}
		}
	}
	return res
}

func diagonalSum2(mat [][]int) int {
	res, n := 0, len(mat)
	for i := 0; i < n; i++ {
		res += mat[i][i]
		res += mat[i][n-1-i]
	}
	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}
