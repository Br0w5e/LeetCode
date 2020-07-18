package main

import "fmt"

//朴素暴力
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}
//二分查找
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left+(right-left)/2
		if nums[mid] < target {
			left = mid+1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid-1
		}
	}
	return left
}

func main()  {
	nums := []int{1,3,5,6}
	target := 2
	fmt.Println(searchInsert1(nums, target))
}