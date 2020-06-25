package main

func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 {
		return maxSide
	}
	rows, columns := len(matrix) , len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				curMaxSide := min(rows-i, columns-j)
				//边长为k
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					//确定边长为k的里边没有0
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide*maxSide
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maximalSquare2(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	columns := len(matrix[0])
	if columns == 0 {
		return 0
	}
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for j := 0; j < columns; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	max := 0
	for _, row := range dp {
		for _, num := range row {
			if num > max {
				max = num
			}
		}
	}
	return max*max
}