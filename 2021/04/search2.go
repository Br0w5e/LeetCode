//查找旋转有序数组中的目标数
//方法一：按序查找
//方法二：二分查找

//33. 搜索旋转排序数组
//注意和81题的区别

package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		//前半部分有序
		if nums[left] <= nums[mid] {
			//前半部分查找
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			// 后半部分有序
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
