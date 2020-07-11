package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	//dp[i][0]: 手上持有股票的最大收益
	//dp[i][1]: 手上不持有股票，且处于冷冻期的最大收益
	//dp[i][2]: 手上不持有股票，未处于冷冻期
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) //购买或者继续上一个
		dp[i][1] = dp[i-1][0] + prices[i]                //卖掉手中有的股票
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])           //继续不持有或者从处于冷冻期的状态到继续不持有
	}
	return max(dp[n-1][1], dp[n-1][2]) //最后肯定不持有的
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
