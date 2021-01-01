package main

//1030. 距离顺序排列矩阵单元格
import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	for i := 0; i < R*C; i++ {
		res[i] = make([]int, 2)
		res[i][0] = i/C
		res[i][1] = i%C
	}
	sort.SliceStable(res, func(i, j int) bool {
		d1 := abs(res[i][0]-r0) + abs(res[i][1]-c0)
		d2 := abs(res[j][0]-r0) + abs(res[j][1]-c0)
		return d1 < d2
	})
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
