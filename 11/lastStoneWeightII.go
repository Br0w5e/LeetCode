package main

//1049. 最后一块石头的重量 II
//dp[i][j]：前i个石头里面，且重量不能超过j的，最佳组合的重量：dp[i][j]
//dp[i][j]：前i个石头里面，且重量不能超过j的，最佳组合的重量：dp[i][j]
func lastStoneWeightII(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	w := sum / 2
	dp := make([][]int, len(stones)+1)
	for i := 0; i <= len(stones); i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= len(stones); i++ {
		for j := 1; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= stones[i-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}
		}
	}
	return sum - 2*dp[len(stones)][w]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
