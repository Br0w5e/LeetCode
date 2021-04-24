package main

//1446. 连续字符
//统计最长连续子串的长度
func maxPower(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		tmp := 1
		for i+1 < len(s) && s[i] == s[i+1]{
			tmp++
			i++
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}