package main

//1725. 可以形成最大正方形的矩形数目
func countGoodRectangles(rectangles [][]int) int {
	m := make(map[int]int)
	for _, rectangle := range rectangles {
		m[min(rectangle[0], rectangle[1])]++
	}
	res, max := 0, 0
	for k, v := range m {
		if (k > max)  || (k == max &&  v > res){
			res = v
			max = k
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
