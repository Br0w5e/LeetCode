package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	rk, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
