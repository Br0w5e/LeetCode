package main

//942. 增减字符串匹配

//理解题意
func diStringMatch(S string) []int {
	l, r := 0, len(S)
	res := make([]int, len(S)+1)
	for i := range S {
		if S[i] == 'I' {
			res[i] = l
			l++
		} else {
			res[i] = r
			r--
		}
	}
	res[len(S)] = r

	return res
}
