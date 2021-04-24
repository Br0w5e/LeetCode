package main

//368. 最大整除子集

import "sort"

//dp
func largestDivisibleSubset(nums []int) []int {
	// 升序
	sort.Ints(nums)
	n := len(nums)
	// 特判
	if n == 0 {
		return make([]int, 0, 0)
	}

	res := make([]int, 0, n)
	// 定义一个状态数组，表示每个元素可以和左边的元素组成的最大整除子集
	dp := make([][]int, n)

	// 遍历数组
	for i := 0; i < n; i++ {
		// 用于临时保存i元素左边最大的整除子集
		temp := make([]int, 0)
		// 遍历该元素左边的所有元素
		for j := i - 1; j >= 0; j-- {
			// 如果小的元素可以整除大的，那么大的自然能整除小的，可以整除小元素的子集，自然也能整除大元素，然后取最大的子集
			if nums[i]%nums[j] == 0 && len(dp[j]) > len(temp) {
				temp = dp[j]
			}
		}
		// 实例化[i]，将temp中的子集都添加进来
		dp[i] = make([]int, len(temp), len(temp)+1)
		copy(dp[i], temp)
		// 自然还需要加上当前元素本身
		dp[i] = append(dp[i], nums[i])
		// 答案取最大子集
		if len(dp[i]) > len(res) {
			res = dp[i]
		}
	}
	return res
}
