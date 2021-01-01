package main

//712. 两个字符串的最小ASCII删除和

//dp[i][j]  要想使s1的前i个与s2的前j个相等，最少需要的删除和
//如果s1的第i个字母和s2的第j个字母相等，则 dp[i][j] = dp[i-1][j-1]
//如果s1的第i个字母和s2的第j个字母不相等， 则 dp[i][j] = min(dp[i-1][j] + int(s1[i-1]), dp[i][j-1] + int(s2[j-1]))，即判断删除哪个的代价小一点
func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i < m+1; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j] + int(s1[i-1]), dp[i][j-1] + int(s2[j-1]))
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
