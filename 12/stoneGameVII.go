package main

//5627. 石子游戏 VII

//dp[i][j]=max(sum(stones[i+1:j])−dp[i+1][j],sum(stones[i:j])−dp[i][j−1])

//博弈dp
func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, n)
		sum[i][i] = stones[i]
		for j := i+1; j < n; j++ {
			sum[i][j] = sum[i][j-1] + stones[j]
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = max(sum[i+1][j]-dp[i+1][j], sum[i][j-1]-dp[i][j-1])
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
