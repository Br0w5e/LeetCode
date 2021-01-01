package main

//188. 买卖股票的最佳时机 IV

import "math"

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt32 >> 1
		sell[i] = math.MinInt32 >> 1
	}
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j]-prices[i])
			sell[j] = max(sell[j], buy[j-1]+prices[i])
		}
	}
	return getMax(sell)
}

func getMax(nums []int) int {
	res := 0
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit5(k int, prices []int) int {
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][1] = math.MinInt64
	}
	for _, price := range prices {
		for i := 1; i <= k; i++ {
			//持有
			dp[i][1] = max(dp[i][1], dp[i-1][0]-price)
			//不持有
			dp[i][0] = max(dp[i][0], dp[i][1]+price)
		}
	}
	return dp[k][0]
}
