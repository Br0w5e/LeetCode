//不想交的线
//dp 跟最长子序列一样
func maxUncrossedLines(A []int, B []int) int {
    n, m := len(A), len(B)
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, m+1)
    }
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if A[i] == B[j] {
                dp[i+1][j+1] = dp[i][j] + 1
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