package main

//面试题 01.05. 一次编辑

//计算编辑距离是不是小于等于1
func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n] <= 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


//简单处理
func oneEditAway(first string, second string) bool {
	//更改长短顺序
	if len(first) < len(second) {
		return oneEditAway(second, first)
	}
	if len(first)-len(second) > 1 {
		return false
	}
	if first == second {
		return true
	}
	for i := 0; i < len(second); i++ {
		if first[i] != second[i] {
			//增或者改
			return first[i+1:] == second[i:] || first[i+1:] == second[i+1:]
		}
	}
	return true
}
