package main

//746. 使用最小花费爬楼梯
//dp
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


//优化状态
func minCostClimbingStairs2(cost []int) int {
	first, second := 0, 0
	for i := 2; i <= len(cost); i++ {
		first, second = second, min(second+cost[i-1], first+cost[i-2])
	}
	return second
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}