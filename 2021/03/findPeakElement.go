package main

//162. 寻找峰值

//好好琢磨
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i + 1] {
			return i
		}
	}
	return len(nums) - 1
}

//二分
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid+1
		}
	}
	return l
}