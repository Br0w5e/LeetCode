package main

import "strings"

func countSegments(s string) int {
	words := strings.Split(s, " ")
	res := 0
	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			res++
		}
	}
	return res
}