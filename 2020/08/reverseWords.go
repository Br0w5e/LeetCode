package main

import "strings"

//反转给定的字符串
//方法：两个库函数的使用
//将www baidu com 转化为com baidu www
func reverseWords(s string) string {
	parts := strings.Split(s, " ")
	size := len(parts)
	var rparts[] string
	for i := size-1; i >= 0; i-- {
		if parts[i] != ""{
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, " ")
}

//将www.baidu.com  转化为 com.baidu.www
func reverseWords1(s string) string {
	parts := strings.Split(s, ".")
	size := len(parts)
	var rparts[] string
	for i := size-1; i >= 0; i-- {
		if parts[i] != "" {
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, ".")
}

func reverseWords2(s string) string {
	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		split[i] = reverse(split[i])
	}
	return strings.Join(split, " ")
}
func reverse(s string) string {
	res := make([]byte, 0)
	for i := len(s); i >= 0 ; i-- {
		res = append(res, s[i])
	}
	return string(res)
}


