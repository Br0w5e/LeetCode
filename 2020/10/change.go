package main
//518. 零钱兑换 II
//dp[i][j]: 若只使用 coins 中的前 i 个硬币的，金额 j的凑法种类。
//背包问题
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i-1] {
				dp[i][j] = dp[i][j]+dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][amount]
}

