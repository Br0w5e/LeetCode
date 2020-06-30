//判断字符串是否是回文串，不考虑标点等
package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isNumberOrAlphabet(s[left]) {
			left++
			continue
		}
		if !isNumberOrAlphabet(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isNumberOrAlphabet(u byte) bool {
	return u >= 'a' && u <= 'z' || u >= '0' && u <= '9'
}
/*
func isPalindrome(s string) bool{
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right{
		//只处理数字和字母
		if !isNumOrAlphabet(s[left]) {
			left++
			continue
		}
		if !isNumOrAlphabet(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isNumOrAlphabet(a byte) bool {
    return ('a' <= a && a <= 'z' || 'A' <= a && a <= 'Z' || '0' <= a && a <= '9')
}
*/