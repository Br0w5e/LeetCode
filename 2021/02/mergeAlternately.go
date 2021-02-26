package main

//5685. 交替合并字符串

//模拟
func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, 0)
	minLen := min(len(word1), len(word2))
	for i := 0; i < minLen; i++{
		res = append(res, word1[i])
		res = append(res, word2[i])
	}
	if minLen == len(word1) {
		for i := minLen; i < len(word2); i++ {
			res = append(res, word2[i])
		}
	} else {
		for i := minLen; i < len(word1); i++ {
			res = append(res, word1[i])
		}
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
