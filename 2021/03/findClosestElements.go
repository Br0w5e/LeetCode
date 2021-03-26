package main

//658. 找到 K 个最接近的元素

//二分查找
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := left+ (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid+1
		} else {
			right = mid
		}
	}
	return arr[left: left+k]
}
