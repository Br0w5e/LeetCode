package main
//486
func PredictTheWinner(nums []int) bool {
	length := len(nums)
	if length%2 == 0 {
		return true
	}
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = nums[i]
	}

	for i := length-2; i >= 0; i-- {
		for j := i+1; j < length; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[length-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//2维dp
func PredictTheWinner2(nums []int) bool {
	length := len(nums)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(nums[i] - dp[i + 1][j], nums[j] - dp[i][j - 1])
		}
	}
	return dp[0][length - 1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//递归
func PredictTheWinner3(nums []int) bool {
	return total(nums, 0, len(nums) - 1, 1) >= 0
}

func total(nums []int, start, end int, turn int) int {
	if start == end {
		return nums[start] * turn
	}
	scoreStart := nums[start] * turn + total(nums, start + 1, end, -turn)
	scoreEnd := nums[end] * turn + total(nums, start, end - 1, -turn)
	return max(scoreStart * turn, scoreEnd * turn) * turn
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}