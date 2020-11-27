package main

//454. 四数相加 II

//map记录
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	m := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			res += m[-(c+d)]
		}
	}
	return res
}
