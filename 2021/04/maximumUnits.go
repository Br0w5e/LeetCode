package main

import "sort"


//1710. 卡车上的最大单元数

//先按照单元数量大的排序，然后进行贪心
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	i := 0
	res := 0
	for truckSize > 0 && i < len(boxTypes) {
		if boxTypes[i][0] < truckSize {
			res += boxTypes[i][0] * boxTypes[i][1]
		} else {
			res += boxTypes[i][1] * truckSize
		}
		truckSize -= boxTypes[i][0]
		i++
	}
	return res
}
