package main

//1353. 最多可以参加的会议数目

//模拟
import "sort"

//暴力--->按照结束时间排序，结束时间相同的时候按照开始时间排序
//超时
func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		return (events[i][1] < events[j][1]) || (events[i][1] == events[j][1] && events[i][0] < events[j][0])
	})

	vis := make([]bool, 100000+1)
	res := 0
	for _, event := range events {
		//遍历可以操作的天数
		for day := event[0]; day <= event[1]; day++ {
			if !vis[day] {
				vis[day] = true
				res++
				break
			}
		}
	}
	return res
}

//记录上一个值
func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	vis, res := make([]bool, 100001), 0
	preStart, preChoice := 0, 0

	for _, event := range events {
		start, end := event[0], event[1]
		if start >= preStart {
			start = max(preChoice, start)
		}

		for i := start; i <= end; i++ {
			if !vis[i] {
				vis[i] = true
				res++

				preStart = event[0]
				preChoice = i + 1

				break
			}
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
