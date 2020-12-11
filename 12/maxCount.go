package main

//598. 范围求和 II

//模拟或者脑筋急转弯
//每次都是从右上角开始的
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m*n
	}
	x, y := 100000, 100000
	for _, op := range ops {
		x = min(x, op[0])
		y = min(y, op[1])
	}
	return x*y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
