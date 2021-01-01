package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 366)
	j := 0
	for i := 1; i < 366; i++ {
		if j >= len(days) {
			//超出最后一次出行
			break
		} else if i < days[j]{
			//不需要出行
			dp[i] = dp[i-1]
		} else if i == days[j] {
			//需要出行
			d7 := 0
			if i > 7 {
				d7 = dp[i-7]
			}
			d30 := 0
			if i > 30 {
				d30 = dp[i-30]
			}
			dp[i] = min(dp[i-1]+costs[0], min(d7+costs[1], d30+costs[2]))
			j++
		}
	}
	return dp[days[j-1]]
}

func mincostTickets2(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	dp[0] = 0
	for k, v := range days {
		if k != 0 {
			for i := days[k-1]+1; i < v; i++ {
				dp[i] = dp[days[k-1]]
			}
		}
		dp[v] = min(dp[max(v-1, 0)]+costs[0], min(dp[max(v-7, 0)]+costs[1], dp[max(v-30, 0)]+costs[2]))
	}
	return dp[days[len(days)-1]]
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

func main()  {
	days := []int{1,2,3,4,5,6,7,8,9,10,30,31}
	costs := []int{2,7,15}
	fmt.Println(mincostTickets2(days, costs))
}