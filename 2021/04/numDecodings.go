package main

//91. 解码方法

//动态规划
//dp[i]：s[:i] 的可用解码方式
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		///判断截取一个是否符合（只要不是0，都符合
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		//判断截取两个是否符合
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10 + (s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

//状态空间优化
func numDecodings(s string) int {
	n := len(s)
	// a = dp[i-2], b = dp[i-1], c = dp[i]
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10 + (s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}