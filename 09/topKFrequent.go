package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	//生成map
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	//寻找最大的k个数字
	tmp := make([]int, 0)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	sort.Ints(tmp)
	res := make([]int, 0)
	for i := len(tmp)-1; i >= len(tmp)-k; i-- {
		for key, v := range m {
			if v == tmp[i] {
				res = append(res, key)
				//标记为0
				m[key] = 0
				break
			}
		}
	}
	return res
}
