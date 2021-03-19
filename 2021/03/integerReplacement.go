package main

import (
	"math"
)

//397. 整数替换
//递归方法
func integerReplacement(n int) int {
	//递归出口
	if n < 3 {
		return n - 1
	}
	if n == math.MinInt32 {
		return 32
	}
	//2的倍数
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}
	//不是2的倍数
	add := integerReplacement(n+1) + 1
	min := integerReplacement(n-1) + 1
	if add > min {
		return min
	}
	return add
}

//非递归方法
func integerReplacement(n int) int {
	//最后两位都是1，且该数不是3是，就选择+1，否则就选择－1
	res := 0
	for n != 1 {
		if n == 3 {
			res += 2
			break
		}
		if n&1 == 0 {
			n = n >> 1
		} else {
			if n&0x3 == 0x3 {
				n++
			} else {
				n--
			}
		}
		res++
	}
	return res
}

func integerReplacement(n int) int {
	res := 0
	for n != 1 {
		if n == 3 {
			res += 2
			break
		}
		if n&1 == 0 {
			n >>= 1
		} else if n&0x3 == 0x3 {
			n++
		} else {
			n--
		}
		res++
	}
	return res
}
