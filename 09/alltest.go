package main

import "fmt"

//此文件为代码测试文件

func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	for i := 0; i < n-k; i++ {
		nums = append(nums, nums[0])
		nums = nums[1:]
	}
	return nums
}

func main() {
	nums := []int{-1, -100, 3, 99}
	fmt.Println(rotate(nums, 2))
}
