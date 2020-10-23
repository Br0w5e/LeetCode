package main

//1619. 删除某些元素后的数组均值
import "sort"

func trimMean(arr []int) float64 {
	left, right := len(arr)/20, len(arr)-len(arr)/20
	sort.Ints(arr)
	arr = arr[left:right]
	sum, mean := 0, 0.0
	for _, num := range arr {
		sum += num
	}
	mean = float64(sum)/float64(len(arr))
	return mean
}