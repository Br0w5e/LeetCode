package main

import "math"

func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n + 1)
	sub := make([]int, n + 1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m + 1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		sub[i + 1] = sub[i] + nums[i]
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j - 1], sub[i] - sub[k]))
			}
		}
	}
	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}