package main

//LCP 28. 采购方案
import "sort"

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			res += (right - left)
			left++
		}
	}
	return res % 1000000007
}
