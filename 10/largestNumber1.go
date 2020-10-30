package main

import "strconv"

//1449. 数位成本和为目标值的最大数字
//dp，不是很懂，记录一下

func largestNumber(cost []int, target int) string {
	dp := make([]string, target+1)
	for i := 1; i <= target; i++ {
		for j := 1; j <= 9; j++ {
			if i-cost[j-1] < 0 {
				continue
			}
			if i-cost[j-1] == 0 {
				dp[i] = max(dp[i], strconv.Itoa(j))
			}
			if i-cost[j-1] > 0 && len(dp[i-cost[j-1]]) > 0 {
				dp[i] = max(dp[i], strconv.Itoa(j)+dp[i-cost[j-1]])
			}
		}
	}
	// fmt.Println(dp)
	if len(dp[target]) == 0 {
		return "0"
	}
	return dp[target]
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	if len(a) < len(b) {
		return b
	}
	if a > b {
		return a
	}
	return b
}
