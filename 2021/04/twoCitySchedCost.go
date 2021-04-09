package main

import "sort"

func twoCitySchedCost(costs [][]int) int {
	//先加上全部去A市
	res := 0
	//去A市比去B市多花的钱
	nums := make([]int, 0)
	for _, cost := range costs {
		res += cost[0]
		nums = append(nums, cost[0]-cost[1])
	}
	sort.Ints(nums)
	//必须有一半要去B市， 选择差额最大的一半
	for i := len(nums) - 1; i >= len(nums)/2; i-- {
		res -= nums[i]
	}
	return res
}
