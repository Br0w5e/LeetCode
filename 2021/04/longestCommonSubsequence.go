package main

//1143. 最长公共子序列


//最长子序列的长度
//DP
//dp[i][j] 表示text1[0~i-1]和text2[0~j-1]的最长公共子序列长度
//相等dp[i+1][j+1]=dp[i][j]+1
//不相等dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
func longestCommonSubsequence(text1 string, text2 string) int {
    n, m := len(text1), len(text2)
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++{
        dp[i] = make([]int, m+1)
    }
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if text1[i] == text2[j] {
                dp[i+1][j+1] = dp[i][j]+1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[n][m]
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}