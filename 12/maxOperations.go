package main

import (
	"fmt"
	"sort"
)

//双指针
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	res := 0
	for left < right {
		if nums[left] + nums[right] == k {
			res++
			left++
			right--
		} else if nums[left] + nums[right] > k {
			right--
		} else {
			left++
		}
	}
	return res
}

//哈希表
func maxOperations(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[k-nums[i]] > 0 {
			res++
			m[k-nums[i]]--
		} else {
			m[nums[i]]++
		}
	}
	return res
}

func main()  {
	nums := []int{4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4}
	k := 2
	fmt.Println(maxOperations(nums, k))
}
