package main

//5605. 检查两个字符串数组是否相等
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	s2 := ""
	for _, word := range word1 {
		s1 += word
	}
	for _,word := range word2 {
		s2 += word
	}
	return s1 == s2
}
