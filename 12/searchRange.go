package main

import "sort"

func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}


//二分
func searchRange2(nums []int, target int) []int {
	n := len(nums)
	res := []int{-1, -1}
	if n == 0 {
		return res
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left) >> 1
		if nums[mid] < target {
			left = mid+1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return res
	}
	res[0] = left
	right = n-1
	for left < right {
		mid := left+1 + (right-left) >> 1
		if nums[mid] > target {
			right = mid-1
		} else {
			left = mid
		}
	}
	res[1] = right
	return res
}

//普遍
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res[0] = i
			break
		}
	}

	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] == target {
			res[1] = i
			break
		}
	}
	return res
}