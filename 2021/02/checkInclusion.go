package main

//567. 字符串的排列

//map记录
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1 := make(map[byte]int, 0)
	m2 := make(map[byte]int, 0)
	n := len(s1)
	for i := 0; i < n; i++ {
		m1[s1[i]]++
		m2[s2[i]]++
	}
	for i := n; i < len(s2); i++ {
		if judge(m1, m2) {
			return true
		}
		m2[s2[i-n]]--
		m2[s2[i]]++
	}
	return judge(m1, m2)
}

func judge(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

//双指针

func checkInclusion2(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	n := len(s1)
	m := make(map[rune]int)
	for _, ch := range s1 {
		m[ch]--
	}
	left := 0
	for right, ch := range s2 {
		m[ch]++
		for m[ch] > 0 {
			m[rune(s2[left])]--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}