package main

//1331. 数组序号转换

//map+排序
import "sort"

func arrayRankTransform(arr []int) []int {
	//拷贝
	arrcp := make([]int, len(arr))
	copy(arrcp, arr)
	//排序
	sort.Ints(arrcp)
	m := make(map[int]int)

	//标记
	index := 1
	for i := 0; i < len(arrcp); i++ {
		if i != 0 && arrcp[i-1] == arrcp[i] {
			continue
		}
		m[arrcp[i]] = index
		index++
	}
	//结果
	res := make([]int, len(arr))
	for i, num := range arr {
		res[i] = m[num]
	}
	return res

}
