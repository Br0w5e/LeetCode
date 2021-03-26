package main

//5693. 字符串中第二大的数字

//模拟
import "sort"

func secondHighest(s string) int {
	nums := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nums = append(nums, int(s[i]-'0'))
		}
	}
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	res := nums[len(nums)-1]
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] < res {
			return nums[i]
		}
	}
	return -1
}
