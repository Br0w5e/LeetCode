package main

//5726. 数组元素积的符号
//统计负数个数
func arraySign(nums []int) int {
	flag := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			flag++
		}
	}
	if flag%2 == 0 {
		return 1
	}
	return -1
}
