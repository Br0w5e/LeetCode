package main

//949. 给定数字能组成的最大时间
import (
	"sort"
	"strconv"
)

//排序+遍历
func largestTimeFromDigits(A []int) string {
	//先排序
	sort.Ints(A)
	for i := 3; i >= 0; i-- {
		//第一位
		if A[i] > 2 {
			continue
		}
		//d第二位
		for j := 3; j >= 0; j-- {
			if j == i || A[i] == 2 && A[j] > 3 {
				continue
			}
			//第三位
			for k := 3; k >= 0; k-- {
				if k == i || k == j || A[k] > 5 {
					continue
				}
				//0+1+2+3，总共4位，和为6
				return strconv.Itoa(A[i]) + strconv.Itoa(A[j]) + ":" + strconv.Itoa(A[k]) + strconv.Itoa(A[6-i-j-k])
			}
		}
	}
	return ""
}
