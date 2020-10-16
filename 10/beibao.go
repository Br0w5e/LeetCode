package main

import "fmt"

//n:物品种类
//c:背包容量
func solve(n int, c int, w []int, v []int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, c+1)
		dp[i][0] = 0
	}
	for i := 0; i < c+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= n ; i++ {
		for j := 1; j <= c; j++ {
			dp[i][j] = dp[i-1][j]
			//还可以装下w[i]
			if j >= w[i-1] {
				dp[i][j] = max(dp[i][j], dp[i][j-w[i-1]]+v[i-1])
			}
		}
	}
	//在1~n这n个物品中挑选若干物品，总质量不超过c，此时能取得的最大价值
	return dp[n][c]
}

func max(a, b int) int  {
	if a > b {
		return a
	}
	return b
}

func main()  {
	w := []int{5, 10, 15}
	v := []int{50, 80, 45}
	fmt.Println(solve(3,25, w, v))
}
