package main

//当面积模宽不等于0时候，一直进行处理，直到宽为1
import "math"

//492. 构造矩形
func constructRectangle(area int) []int {
	W := int(math.Sqrt(float64(area)))
	for area % W != 0 {
		W--
	}
	return []int{area/W, W}
}
