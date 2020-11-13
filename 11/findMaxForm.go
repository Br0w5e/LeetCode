package main

//474. 一和零

//dp[i][j]:使用 i 个 0 和 j 个 1，最多能拼出的字符串数目，隐含遍历到底k个字符串

// 一些说明
// dp[i][j] = dp[i-1][j-1]+1
// 如果采用从后向前填表，那么我们的dp[i-1][j-1]应该是上一轮计算的结果，因为这一轮我们还没有更新过这个值
// 但如果采用从前往后填表，那么我们的dp[i-1][j-1]应该是这一轮计算的结果，因为这一轮我们已经更新过这个值
// 但是我们这个二维dp数组是最初的三维dp数组的一个优化，因此，在状态迁移时，我们需要的是上一轮计算的dp[i-1][j-1]
// 这就是为什么我们要从后往前填表了，主要是保留上一轮计算的结果不被覆盖。
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	//当前字符串
	for _, str := range strs {
		zero, one := count(str)

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				// dp[i-zero][j-one] 为上一轮的，也就是新加字符串之前
				dp[i][j] = max(1+dp[i-zero][j-one], dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func count(str string)(int, int) {
	zero, one := 0, 0
	for _, v := range str {
		if v == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
