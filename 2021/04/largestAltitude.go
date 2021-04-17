package main

//1732. 找到最高海拔
func largestAltitude(gain []int) int {
	res := 0
	h := 0
	for _, g := range gain {
		h += g
		if h > res {
			res = h
		}
	}
	return res
}
