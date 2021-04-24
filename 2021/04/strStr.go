package main

import "strings"

//28. 实现 strStr()

//匹配字符串
func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i < m-n+1; i++ {
		if haystack[i:n+i] == needle {
			return i
		}
	}
	return -1
}

//系统API
func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	return strings.Index(haystack, needle)
}
