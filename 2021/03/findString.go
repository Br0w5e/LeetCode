package main

//面试题 10.05. 稀疏数组搜索
func findString(words []string, s string) int {
	for i, word := range words {
		if word == s {
			return i
		}
	}
	return -1
}
