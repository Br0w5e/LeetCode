package main

//164. 最大间距
import "sort"

func maximumGap(nums []int) int {
	sort.Ints(nums)
	if len(nums) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] - nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}
