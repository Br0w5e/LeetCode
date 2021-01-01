package main

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, str := range wordDict {
		m[str] = true
	}
	return dfs(s, m, map[int][]string{}, 0)
}
func dfs(s string, m map[string]bool, record map[int][]string, idx int) []string {
	if str, ok := record[idx]; ok { // 该位置已经处理过则直接返回对应字符串组合
		return str
	}
	res, n := []string{}, len(s)
	for i := idx+1; i <= n; i++ { // 从idx开始找s[idx:i]在m中的单词
		if m[s[idx:i]] {
			if i != n {
				ret := dfs(s, m, record, i) // 递归得到已拆分的单词
				for _, str := range ret {
					res = append(res, s[idx:i]+" "+str)
				}
			} else {
				res = append(res, s[idx:i])
			}
		}
	}
	record[idx] = res
	return res
}


func wordBreak(s string, wordDict []string) []string {
	if !isValid(s, wordDict) {
		return nil
	}
	n := len(s)
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if len(word) <= i && word == s[i-len(word):i] {
				for _, str := range dp[i-len(word)] {
					if str != "" {
						str += " "
					}
					dp[i] = append(dp[i], str+word)
				}
			}
		}
	}
	return dp[n]
}

func isValid(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if len(word) <= i && word == s[i-len(word):i] && dp[i-len(word)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}