package main

//396. 旋转函数

import "math"

//暴力模拟--->超时
func maxRotateFunction(nums []int) int {
	res := math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		tmp := 0
		for k := 0; k < n; k++ {
			tmp += k * nums[(i+k)%n]
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

//递推关系
//F(n) = 0*A[0]+1*A[1]+...+(n)*A[n];
//F(n-1) = 0*A[1]+1*A[2]+2*A[3]+...+(n-1)*A[n]+n*A[0];
//递推关系：F(i+1) - F(i) = sumA - nA[n-i-1]
func maxRotateFunction(nums []int) int {
	F, sum := 0, 0
	n := len(nums)
	for i, num := range nums {
		sum += num
		F += i * num
	}
	res := F
	for i := 1; i < len(nums); i++ {
		F += sum - n*nums[n-i]
		if F > res {
			res = F
		}
	}
	return res
}
