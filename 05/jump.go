package main
//參考一下canJump
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return  0
	}
	reach, step := 0, 0
	nextReach := nums[0]
	for i := 0; i < n; i++ {
		nextReach = max(nextReach, i+nums[i])
		if nextReach >= n-1 {
			return step+1
		}
		if i == reach {
			step++
			reach = nextReach
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
