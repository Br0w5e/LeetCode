package main
//停留在原地方的方案数目
//动态规划
func numWays(steps int, arrLen int) int {
	MOD := int(1e9+7)
	dp := make([][]int, steps)
	if steps < arrLen {
		arrLen = steps
	}
	for i := 0; i < steps; i++ {
		dp[i]= make([]int, arrLen+2)
	}
	dp[0][1] = 1
	for i := 1; i < steps; i++ {
		for j := 1; j <= arrLen; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1] + dp[i-1][j+1]
			dp[i][j] = dp[i][j] % MOD
		}
	}
	return (dp[steps-1][1] + dp[steps-1][2])%MOD
}
