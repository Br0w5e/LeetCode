package main

//330. 按要求补齐数组

//贪心
func minPatches(nums []int, n int) int {
	res := 0
	x := 0
	for i := 1; i <= n; {
		if x < len(nums) && nums[x] <= i {
			i += nums[x]
			x++
		} else {
			i *= 2
			res++
		}
	}
	return res
}
