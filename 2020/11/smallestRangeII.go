package main

import "sort"
// 910. 最小差值 II
//首先对数组排序，然后最重要的一个思想是，排序后的数组中，一定有一个位置 i，
//i 本身及左测全部加 K,i 右侧全部减 K。 即A[0]..A[i]全部加上K，A[i+1]..A[n-1]全部减去 K， 此时整个数组的最大值是 A[n-1]-K 或 A[i]+K，最小值是 A[0]+K 或 A[i+1]-K
func smallestRangeII(nums []int, K int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[n-1]-nums[0]
	for i := 1; i < n; i++ {
		//最小数要么是A[0]+K、要么是A[i]-K
		minNum := min(nums[0]+K, nums[i]-K)
		//最大数要么是A[n-1]-K、要么是A[i-1]+K
		maxNum := max(nums[n-1]-K, nums[i-1]+K)
		res = min(maxNum-minNum, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
