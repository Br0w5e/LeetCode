package main

//1664. 生成平衡数组的方案数

//dp正负交替前缀和
//dp[i-1] 表示索引 i 左边部分奇偶元素差值
//dp[n]−dp[i]表示索引 i 右边部分奇偶元素差值的相反数
func waysToMakeFair(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i%2 == 0 {
			dp[i] = dp[i-1] + nums[i-1]
		} else {
			dp[i] = dp[i-1] - nums[i-1]
		}
	}
	res := 0
	for i := 1; i < n+1; i++ {
		if dp[i-1] == dp[n]-dp[i] {
			res++
		}
	}
	return res
}
