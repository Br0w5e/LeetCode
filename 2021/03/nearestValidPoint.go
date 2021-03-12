package main

//5680. 找到最近的有相同 X 或 Y 坐标的点
import "math"

//遍历即可
func nearestValidPoint(x int, y int, points [][]int) int {
	max := math.MaxInt64
	res := -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			if dis(p, x, y) < max {
				max = dis(p, x, y)
				res = i
			}

		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dis(p []int, x, y int) int {
	return abs(p[0]-x)+abs(p[1]-y)
}
