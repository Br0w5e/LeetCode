package main

//119. 杨辉三角 II

//朴素做法
func getRow(n int) []int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	for i := 1; i < n+1;i++{
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i{
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[n]
}

//O(k)
func getRow(n int) []int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = 1
		for j := i-1; j > 0; j-- {
			dp[j] += dp[j-1]
		}
	}
	return dp
}
