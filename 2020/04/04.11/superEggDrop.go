//扔鸡蛋
//有K个鸡蛋，N层楼需要扔多少次才能找到鸡蛋不碎的楼层
//不会，背代码！！

func superEggDrop(K int, N int) int {
	dp := make([]int, K+1)
	res := 0
	for dp[K] < N{
		for i := K; i >= 1; i-- {
			dp[i] += dp[i-1] + 1
		}
		res++
	}
	return res
}

func superEggDrop(K int, N int) int {
    dp := make([][]int, K+1)
    for i := 0; i <= K; i++{
        dp[i] = make([]int, N+1)
    }
    for j := 1; j <= N; j++{
        for i := 1; i <= K; i++{
            dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
            if dp[i][j] >= N {
                return j
            }
        }
    }
    return N
}