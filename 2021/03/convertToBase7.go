package main

//504. 七进制数

import "strconv"

func convertToBase7(num int) string {
	res := ""
	nums := make([]int, 0)
	if num < 0 {
		res = "-"
		num = -num
	}
	if num == 0 {
		return "0"
	}
	for num != 0 {
		nums = append(nums, num%7)
		num /= 7
	}
	for i := len(nums)-1; i >= 0; i-- {
		res += strconv.Itoa(nums[i])
	}
	return res
}
