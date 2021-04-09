package main

//1748. 唯一元素的和
//map 记录
func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := 0
	for k, v := range m {
		if v == 1 {
			res += k
		}
	}
	return res
}
