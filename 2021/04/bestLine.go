package main

//面试题 16.14. 最佳直线

import "math"

func bestLine(points [][]int) []int {
	max := 0
	res := []int{0, 1}
	for i := 0; i < len(points)-max; i++ {
		//记录这个斜率的第二个点
		second := make(map[float32]int)
		//记录这个斜率的个数
		amount := make(map[float32]int)
		o := points[i]
		for j := i + 1; j < len(points); j++ {
			//斜率（默认为最大，如果不在同一y轴则计算）
			slope := float32(math.MaxFloat32)
			if float32(points[j][1]-o[1]) != 0 {
				slope = float32(points[j][0]-o[0]) / float32(points[j][1]-o[1])
			}
			if second[slope] == 0 {
				second[slope] = j
			}
			amount[slope]++
			if amount[slope] > max {
				res = []int{i, second[slope]}
				max = amount[slope]
			}
		}
	}
	return res
}
