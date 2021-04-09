package main

import "math"

//154. 寻找旋转排序数组中的最小值 II
// 参考 153 寻找旋转数组中的最小值

//二分查找
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[right] == nums[right-1] {
			right--
			continue
		}
		mid := left + (right-left)>>1
		//排除相同的
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	math.Log2()
	return nums[left]
}
