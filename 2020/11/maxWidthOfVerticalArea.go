package main

//5540. 两点之间不包含任何点的最宽垂直面积
import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})
	max := 0
	for i := 0; i < len(points)-1; i++ {
		if points[i][0]-points[i+1][0] > max {
			max = points[i][0] - points[i+1][0]
		}
	}
	return max
}