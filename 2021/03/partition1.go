package main


//131. 分割回文串

//dp+回溯
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
		}
	}

	splits := make([]string, 0)
	res := make([][]string, 0)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j+1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return res
}
