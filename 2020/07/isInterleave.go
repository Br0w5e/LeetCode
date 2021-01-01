package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1+n2 != len(s3) {
		return false
	}

	dp := make([][]bool, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 1; i < n1+1; i++ {
		dp[i][0] = dp[i-1][0] && (s1[i-1] == s3[i-1])
	}

	for i := 1; i <= n2; i++ {
		dp[0][i] = dp[0][i-1] && (s2[i-1] == s3[i-1])
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1+j] || dp[i][j-1] && s2[j-1] == s3[i-1+j]
		}
	}

	return dp[n1][n2]
}
//优化
func isInterleave1(s1 string, s2 string, s3 string) bool {
	n1, n2 := len(s1), len(s2)
	if (n1 + n2) != len(s3) {
		return false
	}
	dp := make([]bool, n2+1)
	dp[0] = true
	for i := 0; i < n1+1; i++ {
		for j := 0; j < n2+1; j++ {
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[n2]
}

func main()  {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
