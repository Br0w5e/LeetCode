package main
//122. 买卖股票的最佳时机 II
//买卖股票的最大收益
//一次遍历
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
