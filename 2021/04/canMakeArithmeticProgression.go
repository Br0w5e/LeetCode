package main

import "sort"

//1502. 判断能否形成等差数列

//先排序，后算差
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != d {
			return false
		}
	}
	return true
}
