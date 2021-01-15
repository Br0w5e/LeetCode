package main
//1504. 统计全 1 子矩形

//dp

func numSubmat(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][0] = mat[i][0]
	}
	//预处理，dp[i][j]表示向左延伸的1的个数
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if mat[i][j] == 1 {
				dp[i][j] = dp[i][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := dp[i][j]
			//枚举高
			for k := i; k >= 0 && tmp != 0; k-- {
				tmp = min(tmp, dp[k][j])
				res += tmp
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
