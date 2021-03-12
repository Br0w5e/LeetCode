package main

//132. 分割回文串 II
func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	//dp[i][j] 表示i到j为回文串， 当i大于j时为空串（为回文串）
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	//当且仅当dp[i+1][j-1]为回文串 true， 并且s[i] == s[j]时 dp[i][j]才是true
	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && (dp[i+1][j-1])
		}
	}
	//res[i] 表示前缀s[0...i]最少分割的次数
	res := make([]int, n)

	for i := range res {
		//如果s[0...i]为回文串则不用分割 即res[i] = 0
		if dp[0][i] {
			continue
		}
		//做最多有2000个字符
		res[i] = 1999
		for j := 0; j < i; j++ {
			if dp[j+1][i] && res[j]+1 < res[i] {
				res[i] = res[j]+1
			}
		}
	}

	return res[n-1]
}