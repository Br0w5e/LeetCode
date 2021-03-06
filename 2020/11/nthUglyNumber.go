package main
//丑数

//剑指 Offer 49. 丑数
//264. 丑数 II

// dp
// 解题思路：
// 底：dp[0] = 0，dp[1] = 1
// dp[n] = min(dp[a]*2, dp[b]*3, dp[c]*5)
func nthUglyNumber(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	// 有底向上dp
	for a, b, c, i:= 1, 1, 1, 2; i <= n; i++ {
		//选择当前数组中最小的一个
		dp[i] = min(dp[a]*2, min(dp[b]*3, dp[c]*5))
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}