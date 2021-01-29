package main

//674. 最长连续递增序列
//方法一：数组记录，记录每个，最后求最大值返回
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	return getMax(res)
}

func getMax(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return nums[left]
}

//方法二：数组记录，在求解过程中处理最大值
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	max := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			res[i] = res[i-1] + 1
		}
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

//单个变量，不断优化
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, tmp := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			tmp++
		} else {
			tmp = 1
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
