package main

import "sort"

//1046. 最后一块石头的重量
//模拟
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		tmp := stones[n-1]-stones[n-2]
		stones = stones[:n-2]
		if tmp != 0 {
			stones = append(stones, tmp)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}