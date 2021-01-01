package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	n, res := len(A), 0
	for i := 0; i < n; i++{
		for j := i+1; j < n+1; j++ {
			if judge(A[i:j], K) {
				res++
			}
		}
	}
	return res
}

func judge(nums []int, K int) bool {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res%K == 0
}
// 上面的方法超时


//有利用到同余的性质
func subarraysDivByK2(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum % K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}