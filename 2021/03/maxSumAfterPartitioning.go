package main

//1043. 分隔数组以得到最大和

//dp[i] , 代表前i个数分割出来的最大值
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		maxNum := -1
		for j := i - 1; j >= max(i-k, 0); j-- {
			maxNum = max(maxNum, arr[j])
			dp[i] = max(dp[i], dp[j]+maxNum*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
