package main

//213. 打家劫舍 II

//首位成环的，分两种情况称为rob1
//dp
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robber(nums[:n-1]), robber(nums[1:]))
}

func robber(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}
//优化空间
func robber(nums []int) int {
	n := len(nums)
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
