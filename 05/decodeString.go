package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getString(sub []string) string {
	ret := ""
	for _, s := range sub{
		ret += s
	}
	return ret
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}


//递归
var (
	src string
	ptr int
)

func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}

func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits()
		ptr++
		str := getString()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigits() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret * 10 + int(src[ptr] - '0')
	}
	return ret
}