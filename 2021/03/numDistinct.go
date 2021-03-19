package main

//115. 不同的子序列

//动态规划
//dp[i][j] 表示在s[:i]的子序列中t[:j]出现的个数
//当 S[j] == T[i] , dp[i][j] = dp[i-1][j-1] + dp[i][j-1];
//当 S[j] != T[i] , dp[i][j] = dp[i][j-1]
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if j > i {
				continue
			}
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
