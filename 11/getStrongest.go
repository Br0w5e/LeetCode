package main

//1471. 数组中的 k 个最强值
import "sort"

//先排序，确定中位数，然后双指针遍历
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	mid := arr[(len(arr)-1)/2]
	low, high := 0, len(arr)-1
	res := make([]int, k)
	for i := 0; i < k; i++ {
		if abs(arr[low], mid) > abs(arr[high], mid) {
			res[i] = arr[low]
			low++
		} else {
			res[i] = arr[high]
			high--
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}
