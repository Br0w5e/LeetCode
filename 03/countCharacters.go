//拼写单词，给定单词表和字符串进行组合判断
func countCharacters(words []string, chars string) int {
	chars1 := [26]int{}
	for _, ch := range chars {
		chars1[ch-'a']++
	}

	length := 0
	for _, word := range words {
		chars2, tmpLen := chars1, 0
		for _, letter := range word {
			if chars2[letter-'a'] <= 0 {
				tmpLen = 0
				break
			}
			chars2[letter-'a']--
			tmpLen++
		}
		length += tmpLen
	}
	return length
}