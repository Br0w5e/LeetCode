package main

//873. 最长的斐波那契子序列的长度
//dp[i][j]表示以A[i]、A[j]结尾的子序列的最大斐波那契数列长度
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 2
		}
	}
	res := 0
	for i := 1; i < n; i++ {
		l, r := 0, i-1
		for l < r {
			sum := arr[l]+arr[r]
			if sum == arr[i] {
				dp[r][i] = max(dp[r][i], dp[l][r]+1)
				res = max(dp[r][i], res)
				l++
				r--
			} else if sum < arr[i] {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
