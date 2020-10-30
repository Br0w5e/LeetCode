package main

//1464. 数组中两元素的最大乘积
import "sort"

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	return (nums[len(nums)-1]-1)*(nums[len(nums)-2]-1)
}


func maxProduct(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if (nums[i]-1)*(nums[j]-1) > max {
				max = (nums[i]-1)*(nums[j]-1)
			}
		}
	}
	return max
}