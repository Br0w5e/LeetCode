package main

//1287. 有序数组中出现次数超过25%的元素

//统计
func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}
	for k, v := range m {
		if v > len(arr)/4 {
			return k
		}
	}
	return -1
}

//直接判断1/4长度
func findSpecialInteger1(arr []int) int {
	n := len(arr)
	cross := n >> 2
	for i := 0; ; i++ {
		if arr[i] == arr[i+cross] {
			return arr[i]
		}
	}
}
