package main

//杨辉三角
func generate(numRows int) [][]int {
	dp := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		dp = append(dp, tmp)
		dp[i][0] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
		dp[i][i] = 1
	}
	return dp
}
