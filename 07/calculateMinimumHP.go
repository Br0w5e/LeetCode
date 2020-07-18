package main

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := m-1; i >= 0 ; i-- {
		for j := n-1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				//公主位置
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				//最后一行
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				//最后一列
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				//其他
				dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}