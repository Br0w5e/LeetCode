package main

import "sort"

//539. 最小时间差

//字符串转数字，记得头和尾

func findMinDifference(timePoints []string) int {
	//肯定存在重复
	if len(timePoints) > 24*60 {
		return 0
	}
	//转换为数字
	timeTable := make([]int, 0)
	for _, timePoint := range timePoints {
		timeTable = append(timeTable, strToInt(timePoint))
	}
	res := 24*60
	sort.Ints(timeTable)
	//计算相邻之间的差
	for i := 1; i < len(timeTable); i++ {
		//相邻相等
		if timeTable[i] == timeTable[i-1] {
			return 0
		}
		if timeTable[i]-timeTable[i-1] < res {
			res = timeTable[i]-timeTable[i-1]
		}
	}
	//最后一个和第一个的差
	maxDiff := timeTable[len(timeTable)-1]-timeTable[0]
	if maxDiff > 12*60 {
		maxDiff = 24*60-maxDiff
	}

	if maxDiff < res {
		res = maxDiff
	}
	return res
}

func strToInt(str string) int {
	h := int(str[0]-'0')*10+int(str[1]-'0')
	m := int(str[3]-'0')*10+int(str[4]-'0')
	return h*60+m
}
