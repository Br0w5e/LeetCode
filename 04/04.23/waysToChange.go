//硬币
//典型DP问题
func waysToChange(n int) int {
	MOD := int(1e9+7)
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for _, c := range coins {
		for i := 0; i < n-c+1; i++ {
			dp[i+c] += dp[i]
		}
	}
	return dp[n]%MOD
}