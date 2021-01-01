package main

//57. 插入区间

//将能进行合并的区间尽量合并，返回最终的区间结果

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if len(newInterval) == 0 {
		return intervals
	}
	index := 0
	lenIntervals := len(intervals)
	//寻找可能的插入点,判断条件区间右端点值小于等于给定的左端点
	for index < lenIntervals && intervals[index][1] < newInterval[0] {
		index++
	}

	tmpI := index
	//合并区间，合并终止条件，区间左端点值大于给定的右端点
	for index < lenIntervals && intervals[index][0] <= newInterval[1] {
		//左端点，同小取小,这个只运行一次
		newInterval[0] = min(newInterval[0], intervals[index][0])
		//右端点，同大取大，这个可能合并多次
		newInterval[1] = max(newInterval[1], intervals[index][1])
		index++
	}
	// intervals[:tempI] 合并前的区间
	// newInterval 合并的区间
	// intervals[i:] 合并后的区间
	intervals = append(intervals[:tmpI], append([][]int{newInterval}, intervals[index:]...)...)

	return intervals
}

func max(a, b int) int  {
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