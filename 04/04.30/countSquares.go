package main

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix) + 1)
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]) + 1)
	}
	num := 0
	for i := 0; i < len(matrix); i++ {
		for k := 0; k < len(matrix[i]); k++ {
			if matrix[i][k] == 1 {
				dp[i+1][k+1] = Min(dp[i][k], dp[i][k+1], dp[i+1][k]) + 1
				num += dp[i+1][k+1]
			}
		}
	}

	return num
}

func Min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	}
	return z
}