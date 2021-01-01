package main

//5539. 按照频率将数组升序排序

//暴力排序
func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := make([]int, 0)
	for i := 0; i <= 100; i++ {
		for j := 100; j >= -100; j-- {
			if m[j] == i {
				for k := 0; k < i; k++ {
					res = append(res, j)
				}
			}
		}
	}
	return res
}
