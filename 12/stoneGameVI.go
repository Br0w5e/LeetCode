package main

import "sort"

//5611. 石子游戏 VI

//第一种方案是A取第一个，B取第二个，A与B的价值差是 c1 = a1 - b2
//第二种方案是A取第二个，B取第一个，A与B的价值差是 c2 = a2 - b1
//那么这两种方案对于A来说哪一种更优，就取决于两个方案的价值差的比较
//
//记 c = c1 - c2 = （a1 - b2） - (a2 - b1) = (a1 + b1) - (a2 + b2)

//贪心的策略：将两组石头的价值合并，每次取价值最大的那一组。

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	diff := make([][]int, n)
	for i := 0; i < n; i++ {
		diff[i] = make([]int, 2)
		diff[i][0] = i
		diff[i][1] = aliceValues[i]+bobValues[i]
	}
	sort.SliceStable(diff, func(i, j int) bool {
		return diff[i][1] > diff[j][1]
	})
	res := 0
	for i := 0; i + 1 < n; i += 2 {
		res += aliceValues[diff[i][0]]
		res -= bobValues[diff[i+1][0]]
	}
	if n % 2 == 1 {
		res += aliceValues[diff[n-1][0]]
	}
	if res > 0 {
		return 1
	} else if res < 0 {
		return -1
	}
	return 0
}
