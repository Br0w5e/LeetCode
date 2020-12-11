package main

//1509. 三次操作后最大值与最小值的最小差

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	sort.Ints(nums)
	min := math.MaxInt64
	lenth := len(nums)
	if lenth <= 4{
		return 0
	}
	for i:=1;i<=4;i++{
		if nums[lenth-i] - nums[4-i] < min{
			min = nums[lenth-i] - nums[4-i]
		}
	}
	return min
}
