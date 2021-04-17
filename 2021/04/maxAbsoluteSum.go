package main

import (
	"fmt"
	"strings"
)

//1749. 任意子数组和的绝对值的最大值

//前缀和
func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	res := 0
	minNum, maxNum := 0, 0
	for i := 1; i <= n; i++ {
		res = max(res, abs(preSum[i]-minNum))
		res = max(res, abs(preSum[i]-maxNum))
		maxNum = max(maxNum, preSum[i])
		minNum = min(minNum, preSum[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//最大的前缀和减去最小的前缀和
func maxAbsoluteSum(nums []int) int {
	preSum := 0
	min, max := 0, 0
	for _, num := range nums {
		preSum += num
		if preSum > max {
			max = preSum
		}
		if preSum < min {
			min = preSum
		}
	}
	return max - min
	strings.HasPrefix()
}

func main() {

}