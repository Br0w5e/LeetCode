package main

import "sort"


//594. 最长和谐子序列
//为什么能这么做呀？
//为什么能这么做呀？
func findLHS(nums []int) int {
	sort.Ints(nums)
	left, res := 0, 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 1 {
			left++
		}
		if nums[right]-nums[left] == 1 {
			res = max(res, right-left+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := 0
	for k, v := range m {
		if n, ok := m[k-1]; ok {
			if n+v > res {
				res = n+v
			}
		}

		if n, ok := m[k+1]; ok {
			if n+v > res {
				res = n+v
			}
		}
	}
	return res
}