package main

//1005. K 次取反后最大化的数组和

//两次排序
import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	//排序
	sort.Ints(A)
	//负数变号，直到K=0,或者没有负数
	for i := 0; i < len(A) && K > 0; i++ {
		if A[i] > 0 {
			break
		} else {
			A[i] = -A[i]
		}
		K--
	}
	//重新排序
	sort.Ints(A)
	res := sum(A)
	//当K = 0 或者 K为偶数的情况
	if K%2 == 0 {
		return res
	}
	//当K为奇数的情况
	return res - 2*A[0]
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
