package main

import "sort"

//628. 三个数的最大乘积

//排序、要么最后三个的乘积，要么前两个和最后一个的乘积。
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[0]*nums[1]*nums[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
