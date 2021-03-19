package main

//3. 无重复字符的最长子串
//剑指 Offer 48. 最长不含重复字符的子字符串
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

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]int)
	left := 0
	for right := 0; right < len(s); right++ {
		//大于0
		if m[s[right]] > 0 {
			for m[s[right]] > 0 {
				m[s[left]]--
				left++
			}
			m[s[right]]++
		} else {
			m[s[right]]++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
