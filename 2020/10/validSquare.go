package main

//593. 有效的正方形
//证明四个直角三角形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return isRightTriangle(p1, p2, p3) && isRightTriangle(p1, p2, p4) && isRightTriangle(p1, p3, p4) && isRightTriangle(p2, p3, p4);
}

//勾股定理
func isRightTriangle(p1 []int, p2 []int, p3 []int) bool {
	d1 := (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1])
	d2 := (p2[0] - p3[0]) * (p2[0] - p3[0]) + (p2[1] - p3[1]) * (p2[1] - p3[1])
	d3 := (p3[0] - p1[0]) * (p3[0] - p1[0]) + (p3[1] - p1[1]) * (p3[1] - p1[1])
	if (d1 > d2 && d2 == d3 && d2+d3 == d1) || (d2 > d1 && d1 == d3 && d1 + d3 == d2) || (d3 > d1 && d1 == d2 && d1 + d2 == d3)  {
		return true
	}
	return false
}
