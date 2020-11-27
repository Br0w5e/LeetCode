package main

//452. 用最少数量的箭引爆气球
import "sort"

//交集
func findMinArrowShots(points [][]int) int {
	if len(points) < 1 {
		return 0
	}
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		//有边界不相交
		if right < points[i][0] {
			res++
			right = points[i][1]
		}
	}
	return res
}