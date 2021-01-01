package main

import "strconv"
//青蛙跳台阶
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0 //p 上次循环计算结果，q 本次循环计算结果
		r += q // 最终结果
		if i == 0 {
			continue
		}
		pre := src[i-1:i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
