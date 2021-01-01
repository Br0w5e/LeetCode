package main

import "sort"

//1122. 数组的相对排序

//map
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	sort.Ints(arr1)
	//形成映射表
	for _, num := range arr1 {
		m[num]++
	}
	res := make([]int, 0)
	//处理在arr2中的
	for i := 0; i < len(arr2); i++ {
		for m[arr2[i]] > 0 {
			res = append(res, arr2[i])
			m[arr2[i]]--
		}
	}
	//处理只在arr1中出现的
	for i := 0; i < len(arr1); i++{
		for m[arr1[i]] > 0 {
			res = append(res, arr1[i])
			m[arr1[i]]--
		}
	}
	return res
}

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	//最后一个arr1处理的位置
	index := 0
	for i := 0; i < len(arr2); i++ {
		for j := index; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				arr1[index], arr1[j] = arr1[j], arr1[index]
				index++
			}
		}
	}
	sort.Ints(arr1[index:])
	return arr1
}

