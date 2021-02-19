package main

import "sort"

//1004. 最大连续1的个数 III
//滑动窗口
func longestOnes(A []int, K int) int {
	res := 0
	left, right := 0, 0
	cnt := 0
	for right < len(A) {
		if A[right] == 0 {
			cnt++
		}
		right++

		for cnt > K {
			if A[left] == 0 {
				cnt--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


//更加纯熟的操作
func longestOnes(A []int, K int) int {
	l, r := 0, 0
	for r < len(A) {
		if A[r] == 0 {
			K--
		}
		if K < 0 {
			if A[l] == 0 {
				K++
			}
			l++
		}
		r++
	}
	return r - l
}