package main

//dp[i] 表示跳到第i个位置的最大得分 dp[i] = max(dp[i-1], dp[i-2]...dp[i-k])+nums[i]
//dp ---> 超时
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		maxVal := dp[i-1]
		for j := i - 2; j >= i-k && j >= 0; j-- {
			if dp[j] > maxVal {
				maxVal = dp[j]
			}
			if nums[j] > 0 {
				break
			}
		}
		dp[i] = maxVal + nums[i]
	}
	return dp[len(dp)-1]
}

//dp+单调队列优化
//维护一个不大于k的长度的窗口, 保持里面对应的dp值单调递减. 每次取队头元素
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	queue := make([]int, 1, k)
	for i := 1; i < len(dp); i++ {
		if queue[0] < i-k {
			queue = queue[1:]
		}
		dp[i] = nums[i] + dp[queue[0]]
		//从队尾出队，队里边的元素单调递减的
		for len(queue) > 0 && dp[queue[len(queue)-1]] <= dp[i] {
			queue = queue[:len(queue)-1]
		}
		//将小于的入队
		queue = append(queue, i)
	}
	return dp[len(dp)-1]
}
