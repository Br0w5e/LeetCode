package main

//5718. 统计一个圆中点的数目

//计算距离
func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i, querie := range queries {
		for _, point := range points {
			if judge(querie[0], querie[1], querie[2], point[0], point[1]) {
				res[i]++
			}
		}
	}
	return res
}

func judge(ox, oy, r, x, y int) bool {
	if (x-ox)*(x-ox)+(y-oy)*(y-oy) <= r*r {
		return true
	}
	return false
}
