package main

import "sort"

//896. 单调数列

//一次遍历，标记
func isMonotonic(A []int) bool {
	inc, dec := true, true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			inc = false
		}
		if A[i] < A[i+1] {
			dec = false
		}
	}
	return inc || dec
}

//内置函数
func isMonotonic2(A []int) bool {
	return sort.IntsAreSorted(A) || sort.IsSorted(sort.Reverse(sort.IntSlice(A)))
}