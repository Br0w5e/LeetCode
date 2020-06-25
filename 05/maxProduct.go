package main

//最大子数列成绩
//暴力走一趟
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		tmp := mul(nums[i:])
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func mul(ints []int) int {
	max, tmp := 0, 1
	for i := 0; i < len(ints); i++ {
		tmp *= ints[i]
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func maxProduct2(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}