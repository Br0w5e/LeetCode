package main

func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	//存储所有气球，带有头尾节点的
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	for i := n-1; i >= 0; i-- {
		for j := i+2; j < n+2; j++ {
			for k := i+1; k < j; k++ {
				sum := val[i]*val[k]*val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
