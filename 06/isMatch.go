package main

import "regexp"

func isMatch(s string, p string) bool {
	return len(regexp.MustCompile(p).FindSubmatch([]byte(s))) > 0 && s == string(regexp.MustCompile(p).FindSubmatch([]byte(s))[0])
}
//动态规划
func isMatch2(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 2; i < len(p)+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++{
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}