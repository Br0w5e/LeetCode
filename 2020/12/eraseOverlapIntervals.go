package main

//435. 无重叠区间
import "sort"

//dp
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	maxNum := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxNum = max(maxNum, dp[i])
	}
	return n-maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}