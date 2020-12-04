package main

//525. 连续数组
func findMaxLength(nums []int) int {
	dp := make(map[int]int, len(nums))
	sum := 0
	res := 0
	dp[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		if val, ok := dp[sum]; ok {
			res = max(res, i-val)
		} else {
			dp[sum] = i
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
