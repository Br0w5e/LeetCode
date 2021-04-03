package main

//5713. 字符串中不同整数的数目

//遍历
func numDifferentIntegers(word string) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			w := int(word[i]-'0')
			for i+1 < len(word) && word[i+1] >= '0' && word[i+1] <= '9' {
				i++
				w = w*10 + int(word[i]-'0')
			}
			m[w]++
		}
	}
	return len(m)
}
