package main

//面试题 08.01. 三步问题
//到达台阶的方法
//直接dp
func waysToStep(n int) int {
	mod := int(1e9)+7
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1]%mod+dp[i-2]%mod+dp[i-3]%mod
	}
	return dp[n-1]%mod
}
//优化后的dp
func waysToStep2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	dp1, dp2, dp3 := 1, 2, 4
	for i := 4; i <= n; i++ {
		dp1, dp2, dp3 = dp2, dp3, (dp1+dp2+dp3)%1000000007
	}
	return dp3
}
