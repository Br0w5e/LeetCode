package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closeSum := nums[0]+nums[1]+nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			three := nums[i]+nums[l]+nums[r]
			if abs(three-target) < abs(closeSum-target) {
				closeSum = three
			}
			if three > target {
				r--
			} else if three < target {
				l++
			} else {
				return target
			}
		}
	}
	return closeSum
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

