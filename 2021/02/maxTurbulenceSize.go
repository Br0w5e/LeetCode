package main

//978. 最长湍流子数组

//参考376. 摆动序列， 这个相当于摆动子数组
func maxTurbulenceSize(arr []int) int {
	up, down := 1, 1
	res := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			up = down + 1
			down = 1
		} else if arr[i] < arr[i-1] {
			down = up + 1
			up = 1
		} else {
			up, down = 1, 1
		}
		res = max(res, max(up, down))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
