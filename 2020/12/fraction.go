package main

//LCP 02. 分式化简
func fraction(cont []int) []int {
	res := make([]int, 2)
	res[0] = 1
	res[1] = 0
	for i := len(cont)-1; i >= 0; i--{
		tmp := res[1]
		res[1] = res[0]
		res[0] = cont[i]*res[1] + tmp
	}
	return res
}