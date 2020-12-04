package main
//767. 重构字符串

//贪心+计数
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}
	tmp := make([]int, 26)
	max := 0
	for _, ch := range S {
		tmp[ch-'a']++
		if tmp[ch-'a'] > max {
			max = tmp[ch-'a']
		}
	}

	if max > (n+1) / 2 {
		return ""
	}
	res := make([]byte, n)
	even, odd, half := 0, 1, n/2
	for i, c :=range tmp {
		ch := byte('a'+i)
		for c > 0 && c <= half && odd < n {
			res[odd] = ch
			c--
			odd += 2
		}
		for c > 0 {
			res[even] = ch
			c--
			even += 2
		}
	}
	return string(res)
}
