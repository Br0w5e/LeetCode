package main

//81. 搜索旋转排序数组 II

//思想先处理有序的，无序的放到else中，二分
//注意和31题的区别

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		//查询到结果
		if nums[mid] == target {
			return true
		}

		//相等的情况，关键步骤
		if nums[mid] == nums[left] {
			left++
			continue
		}
		//左半边有序
		if nums[mid] > nums[left] {
			//target在左半边
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
				//target在右半边
			} else {
				left = mid + 1
			}
			//右半部分有序
		} else {
			//target在右半边
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
				//target在左半边
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

//顺序查找
func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}