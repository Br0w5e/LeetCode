package main

//5609. 统计一致字符串的数目

//map标记
func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]bool)
	for _, ch := range allowed {
		m[ch] = true
	}
	res := 0
	for _, word := range words {
		flag := true
		for _, ch := range word {
			if m[ch] == false {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}
