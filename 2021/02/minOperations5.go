package main

//5676. 生成交替二进制字符串的最少操作数

//01串和10串取最小值
func minOperations(s string) int {
	n := len(s)
	cnt1, cnt2 := 0, 0
	for i := 0; i < n; i++ {
		if int(s[i]-'0') == i%2 {
			cnt1++
		} else {
			cnt2++
		}
	}
	if cnt1 < cnt2 {
		return cnt1
	}
	return cnt2
}
