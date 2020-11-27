package main

//686. 重复叠加字符串匹配
import "strings"

func repeatedStringMatch(A string, B string) int {
	la, lb := len(A), len(B)
	maxLen := la * 2 + lb
	t, count := "", 0
	for len(t) < maxLen {
		if strings.Contains(t, B) {
			return count
		}
		t += A
		count++
	}
	return -1
}
