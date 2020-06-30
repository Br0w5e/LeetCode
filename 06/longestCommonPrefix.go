package main

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := " "
	//最短单词长度
	minLen := math.MaxInt32
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		for _, str := range strs {
			if str[i] != strs[0][i] {
				return res[1:len(res)-1]
			} else {
				if str[i] == res[len(res)-1] && i < len(res)-1 {
					continue
				}
				res += string(str[i])
			}
		}
	}
	return res[1:]

}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}



func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	minLen := len(prefix)
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			prefix = str
		}
	}
	for i := 0; i < len(strs); i++ {
		for !hasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}