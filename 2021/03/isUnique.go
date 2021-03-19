package main

//面试题 01.01. 判定字符是否唯一

func isUnique(astr string) bool {
	m := make(map[rune]int)
	for _, ch := range astr {
		m[ch]++
		if m[ch] > 1 {
			return false
		}
	}
	return true
}
