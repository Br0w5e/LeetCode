package main
//714. 买卖股票的最佳时机含手续费
//dp[i][0] 不持有股票最大收益
//dp[i][1] 持有股票的最大收益
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	s, b := 0, -prices[0]
	for i := 1; i < n; i++ {
		s = max(s, b+prices[i]-fee)
		b = max(b, s-prices[i])
	}
	return s
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

//贪心
func maxProfit2(prices []int, fee int) int {
	res := 0
	hold := prices[0]// 记录最低价格
	for i := 1; i < len(prices); i++ {
		if prices[i] < hold {
			// 买入
			hold = prices[i]
		} else if prices[i] > hold+fee {
			// 计算利润，可能有多次计算利润，最后一次计算利润才是真正意义的卖出
			res += prices[i]-fee-hold
			hold = prices[i]-fee
		}
	}
	return res
}
