package main

import "math"

// 400. 第N个数字
// 剑指 Offer 44. 数字序列中某一位的数字

// 1 2 3 4 5 6 7 8 9
// 10...   99 2*90  180
// 100...  999 3*900 2700
// 1000... 9999
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	cur := 0
	bit := 1
	res := 0
	for cur+int(math.Pow(10, float64(bit-1)))*bit*9 < n {
		cur += int(math.Pow(10, float64(bit-1))) * bit * 9
		bit++
	}

	t := int(math.Pow(10, float64(bit-1)))
	// fmt.Println(cur, t, bit, (n-cur)%bit,  (t + (n-cur)/bit), (n-cur)%bit)
	num := t + (n-cur)/bit
	// fmt.Println(num)

	idx := (n - cur) % bit
	if idx == 0 {
		num--
		res = num % 10
	} else if idx > 0 {
		idx = bit - idx
		res = (num / int(math.Pow(10, float64(idx)))) % 10
	}

	return res
}
