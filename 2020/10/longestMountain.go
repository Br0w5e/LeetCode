package main

import (
	"fmt"
	"sort"
)
//845. 数组中的最长山脉
func longestMountain(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	//记录nums[i]可能组成的山峰
	res := make([]int, n)
	for i := 0; i < n; i++ {
		left, right := i, i
		for left > 0 && A[left] > A[left-1] {
			left--
		}
		for right < len(A)-1 && A[right] > A[right+1] {
			right++
		}
		//022不算山峰
		if A[i] != A[right] && A[i] != A[left] {
			res[i] = right-left+1
		}
	}
	//第一个和最后一个不能是peak
	if res[n-1] == n || res[0] == n {
		return 0
	}
	//返回最大的
	sort.Ints(res)
	return res[n-1]
}

func main()  {
	A := []int{0,1,2,3,4,5,4,3,2,1,0}
	fmt.Println(longestMountain(A))
}
