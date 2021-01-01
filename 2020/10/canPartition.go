package main

//平分集合

//416. 分割等和子集
//二维DP
// dp[i][j] 表示从数组的[0,i]下标范围内选取若干个正整数（可以是 0 个），是否存在一种选取方案使得被选取的正整数的和等于 j。

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	max := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum < 2*max || sum%2 == 1 {
		return false
	}
	if sum == 2 * sum {
		return true
	}

	dp := make([][]bool, n)
	target := sum >> 1
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			dp[i][j] = dp[i-1][j]
			////选取不选取当前数字
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}

//一维
//dp[j]=dp[j] ∣ dp[j−nums[i]]
func canPartition1(nums []int) bool {
	n := len(nums)
	//只有一个元素
	if n < 2 {
		return false
	}
	//求和
	max, sum := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	//和为奇数或者最大的数字大于一半
	if sum%2 == 1 || max > sum/2 {
		return false
	}
	target := sum/2
	//最大的等于一半
	if max == sum/2 {
		return true
	}

	//一般情况进行动态规划
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n ; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}