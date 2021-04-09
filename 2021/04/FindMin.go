//寻找最小的，一般方法
package main

//153 寻找旋转数组中的最小值
//参考 154. 寻找旋转排序数组中的最小值 II
//普通方法，直接查找
func findMin(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}
	return nums[0]
}

//二分法
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		//在递增的路上，因此最小值一定在前边
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
