package main

import "sort"

//581. 最短无序连续子数组

//copy+排序+双指针
func findUnsortedSubarray(nums []int) int {
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	sort.Ints(sortNums)
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 && sortNums[p1] == nums[p1]{
		p1++
	}
	for p1 <= p2 && sortNums[p2] == nums[p2] {
		p2--
	}
	return p2-p1+1
}
