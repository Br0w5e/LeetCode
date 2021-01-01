package main

//322. 零钱兑换
//最少的硬币，那么最开始的时候就应该将状态设置为最大的

import "fmt"

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 0
		for j := 1; j < amount+1; j++ {
			dp[i][j] = amount+1
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i][j], dp[i][j - coins[i-1]]+1)
			}
		}
	}
	if dp[n][amount] > amount {
		return -1
	}
	return dp[n][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	coins := []int{1, 2, 5}
	n := 11
	fmt.Println(coinChange(coins, n))
}
