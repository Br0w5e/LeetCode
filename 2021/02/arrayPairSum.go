package main

//561. 数组拆分 I

//先排序，后挨个两两组队
import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i+=2 {
		res += nums[i]
	}
	return res
}
