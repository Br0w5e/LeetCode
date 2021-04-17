package main

//丑数

//剑指 Offer 49. 丑数
//264. 丑数 II

// dp[i] 表示第i个丑数
// 解题思路：
// 底：dp[0] = 0，dp[1] = 1
// dp[n] = min(dp[a]*2, dp[b]*3, dp[c]*5)
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
