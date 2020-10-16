package main

import "sort"

//1552. 两球之间的磁力
//结果导向
func maxDistance(position []int, m int) int {
	//排序
	sort.Ints(position)
	left, right := 0, position[len(position)-1]
	//肯定存在结果的
	res := -1
	//二分
	for left <= right {
		mid := left + (right-left)>>1
		//判断mid情况下，是否可以放至少n个球
		if judge(position, mid, m) {
			res = mid
			//结果有点小，可以进一步扩大
			left = mid + 1
		} else {
			//结果过大，减小
			right = mid - 1
		}
	}
	return res
}

func judge(position []int, mid int, m int) bool {
	pos := position[0]
	//最多可以放球的数目
	cnt := 1
	for _, p := range position {
		if p-pos >= mid {
			//相邻小球的间距均大于等于 mid
			pos = p
			cnt++
		}
	}
	//最多放的球大于题目的要求
	return cnt >= m
}
