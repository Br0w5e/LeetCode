package main

//740. 删除与获得点数

//转化为打家劫舍，要删除的就是相邻的
func deleteAndEarn(nums []int) int {
	bucket := make([]int, 10001)
	for _, num := range nums {
		bucket[num] += num
	}
	dp := make([]int, 10001)
	dp[1] = bucket[1]
	for i := 2; i <= 10000; i++ {
		dp[i] = max(dp[i-2]+bucket[i], dp[i-1])
	}
	return dp[10000]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


//直接dp
func deleteAndEarn(nums []int) int {
	bucket := make([]int, 10001)
	for _, num := range nums {
		bucket[num]++
	}
	pre, cur := 0, 0
	for i := 0; i < 10001; i++ {
		pre, cur = cur, max(bucket[i]*i+pre, cur)
	}
	return cur
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}