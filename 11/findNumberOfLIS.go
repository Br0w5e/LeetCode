package main

//673. 最长递增子序列的个数

//使用DP实现
//假设对于以 nums[i] 结尾的序列，我们知道最长序列的长度 length[i]，以及具有该长度的序列的 count[i]。
//对于每一个 i<j 和一个 A[i]<A[j]，我们可以将一个 A[j] 附加到以 A[i] 结尾的最长子序列上。
//如果这些序列比 length[j] 长，那么我们就知道我们有count[i] 个长度为 length 的序列。如果这些序列的长度与 length[j] 相等，那么我们就知道现在有 count[i] 个额外的序列（即 count[j]+=count[i]）

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	//两个状态变量
	//dp为到达第i个位置时的最长子序列长度
	dp := make([]int, len(nums))
	//count表示到达第I个位置时的最长子序列长度有几种情况
	count := make([]int, len(nums))
	//初始值
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}
	max := 1 //定义初始化最大值
	res := 0 //定义最长递增子序列的个数
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//如果dp[i] < dp[j]+1 表明在0-i-1中，出现了
				//新的最大值，则将最大值更新
				if dp[i] < dp[j]+1 {
					//更新最大值
					dp[i] = dp[j] + 1
					//此时count[i]=count[j]；以nums[i]结尾
					//的最长递增子序列的组合方式就等于nums[i]目前的组合方式
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 { //当相等时，又发现一种情况，
					//将现在的情况与j时的组合方式相加
					count[i] += count[j]
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	//  遍历遍历dp[i]:找到最大长度，
	//然后将结果加上相应的组合长度
	for i := 0; i < len(count); i++ {
		if dp[i] == max {
			res += count[i]
		}
	}
	return res
}
