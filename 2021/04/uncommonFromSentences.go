package main

//884. 两句话中的不常见单词

import "strings"

func uncommonFromSentences(A string, B string) []string {
	mA := make(map[string]int)
	mB := make(map[string]int)
	for _, s := range strings.Split(A, " ") {
		mA[s]++
	}
	for _, s := range strings.Split(B, " ") {
		mB[s]++
	}
	res := make([]string, 0)
	for k, v := range mA {
		if v == 1 && mB[k] == 0 {
			res = append(res, k)
		}
	}

	for k, v := range mB {
		if v == 1 && mA[k] == 0 {
			res = append(res, k)
		}
	}

	return res
}
